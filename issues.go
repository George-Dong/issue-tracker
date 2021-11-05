package main

import (
	"context"
	"encoding/json"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
)

func getIssuesByTimeRange(owner, name string, labels []string, from time.Time, to time.Time, batchLimit int, totalLimit int) (issues []IssueNode, err error) {
	var query struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Cursor githubv4.String
					Node   IssueNode
				}
			} `graphql:"issues(first: $limit, after: $cursor, orderBy: {field: UPDATED_AT, direction: ASC}, labels: $labels, filterBy: {since: $since})"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	cursor := (*githubv4.String)(nil)
	total := 0
	ghLabels := make([]githubv4.String, 0, len(labels))
	for _, l := range labels {
		ghLabels = append(ghLabels, githubv4.String(l))
	}

	since := from.Add(-1 * time.Minute).Format(time.RFC3339)
	log.Printf("fetching since %s", since)

	for totalLimit != 0 {
		limit := batchLimit
		if totalLimit > 0 && totalLimit < limit {
			limit = totalLimit
		}
		param := map[string]interface{}{
			"name":   githubv4.String(name),
			"owner":  githubv4.String(owner),
			"limit":  githubv4.Int(limit),
			"cursor": cursor,
			"labels": ghLabels,
			"since":  githubv4.String(since),
		}

		err = client.Query(context.Background(), &query, param)
		if err != nil {
			log.Println(err)
			return
		}
		edges := query.Repository.Issues.Edges

		for _, edge := range edges {
			issues = append(issues, edge.Node)
			log.Printf("%06d %s %s\n", edge.Node.Number, edge.Node.UpdatedAt.Format(time.RFC3339), edge.Node.Title)
		}

		cnt := len(edges)
		if cnt != 0 {
			lastIssue := &edges[cnt-1]
			cursor = &lastIssue.Cursor
			lastUpdated := lastIssue.Node.UpdatedAt.Time
			total += cnt
			totalLimit -= cnt
			log.Println(cnt, "fetced", owner, name, labels, lastUpdated)
			if lastUpdated.After(to) {
				break
			}
		}
		if cnt != limit {
			break
		}
	}

	log.Printf("fetched %d issues from %s/%s\n", total, owner, name)
	return
}

type IDMap map[githubv4.ID]int

type TrackedIssues struct {
	issues    []IssueNode
	issuesMap IDMap
	closedBy  map[githubv4.ID]githubv4.ID
}

func (ti *TrackedIssues) Load(data []byte) {
	json.Unmarshal(data, &ti.issues)
	log.Printf("load %d issues", len(ti.issues))
	ti.issuesMap = make(IDMap)
	for i, issue := range ti.issues {
		ti.issuesMap[issue.ID] = i
	}
}

func (ti *TrackedIssues) Save() []byte {
	ti.Normalize()
	issuesJson, err := json.MarshalIndent(ti.issues, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return issuesJson
}

func (ti *TrackedIssues) Add(updatedIssues []IssueNode) {
	log.Printf("adding %d issues to %d issues", len(updatedIssues), len(ti.issues))
	for _, issue := range updatedIssues {
		if i, ok := ti.issuesMap[issue.ID]; ok {
			if issue.UpdatedAt.Time.After(ti.issues[i].UpdatedAt.Time) {
				ti.issues[i] = issue
			}
		} else {
			ti.issuesMap[issue.ID] = len(ti.issues)
			ti.issues = append(ti.issues, issue)
		}
	}
	log.Printf("%d issues after adding", len(ti.issues))
}

func (ti *TrackedIssues) Normalize() {
	sort.Slice(ti.issues, func(i, j int) bool {
		return ti.issues[i].UpdatedAt.Time.After(ti.issues[j].UpdatedAt.Time)
	})
	for i, issue := range ti.issues {
		ti.issuesMap[issue.ID] = i
	}
}

func (ti *TrackedIssues) getUpdateTimeRange() (from time.Time, to time.Time) {
	if len(ti.issues) == 0 {
		// if we don't have any tracked things, we initialize from 2 days ago
		from = time.Now().Add(-48 * time.Hour)
		to = from
		return
	}
	from = ti.issues[0].UpdatedAt.Time
	to = from
	for _, issue := range ti.issues {
		if issue.UpdatedAt.Time.After(to) {
			to = issue.UpdatedAt.Time
		}
		if from.After(issue.UpdatedAt.Time) {
			from = issue.UpdatedAt.Time
		}
	}
	return
}

func (ti *TrackedIssues) UpdateByTimeRange(from, to time.Time) (err error, fetched int, added int) {
	var updated []IssueNode
	updated, err = getIssuesByTimeRange("pingcap", "tidb", []string{"type/bug"}, from, to, 20, 500)
	if err != nil {
		log.Printf("error fetching issues %v", err)
		return
	}
	fetched = len(updated)
	original := len(ti.issues)
	ti.Add(updated)
	added = len(ti.issues) - original
	return
}

func (ti *TrackedIssues) Update() {
	var err error
	var updatedIssues []IssueNode

	from, to := ti.getUpdateTimeRange()
	log.Println("issue update time range", from, to)
	labels := []string{"type/bug"}

	updatedIssues, err = getIssuesByTimeRange("pingcap", "tidb", labels, to, time.Now(), 20, 500)
	if err != nil {
		log.Printf("error fetching issues %v", err)
		return
	}
	ti.Add(updatedIssues)

	chunkBy := 48 * time.Hour
	cnt := 0
	earliestTracked := from
	for {
		earlier := earliestTracked.Add(-chunkBy)
		updatedIssues, err = getIssuesByTimeRange("pingcap", "tidb", trackedLabels[0], earlier, earliestTracked, 20, 500)
		if err == nil {
			ti.Add(updatedIssues)
			earliestTracked = earlier
			cnt += len(updatedIssues)
			if cnt > 300 {
				break
			}
		} else {
			log.Println(err)
			break
		}
	}
}

func (ti *TrackedIssues) PopulateClosedBy(tpr *TrackedPullRequests) {
	ti.closedBy = make(map[githubv4.ID]githubv4.ID)
	for _, i := range ti.issues {
		if i.State == githubv4.IssueStateClosed {
			for _, edge := range i.TimelineItems.Edges {
				closer := edge.Node.ClosedEvent.Closer.PullRequest
				if closer.Number != 0 {
					tpr.add(closer)
					ti.closedBy[i.ID] = closer.ID
				}
			}
		}
	}
	log.Printf("populated %d closed by relationship", len(ti.closedBy))
}

var LabelSeverityPrefix = "severity/"
var LabelAffectedVersionPrefix = "affects-"

func GetClosedIssueInfo(t *TrackedIssues, p *TrackedPullRequests) (infos []ClosedIssueInfo) {
	for _, i := range t.issues {
		if i.State == githubv4.IssueStateClosed {
			info := ClosedIssueInfo{}
			info.Title = string(i.Title)
			info.Number = int(i.Number)
			info.Url = string(i.Url)
			info.ClosedAt = i.ClosedAt.Time
			for _, label := range i.Labels.Nodes {
				if strings.HasPrefix(string(label.Name), LabelSeverityPrefix) {
					info.Severity = strings.TrimPrefix(string(label.Name), LabelSeverityPrefix)
				}
				if strings.HasPrefix(string(label.Name), LabelAffectedVersionPrefix) {
					info.AffectedVersions = append(info.AffectedVersions, strings.TrimPrefix(string(label.Name), LabelAffectedVersionPrefix))
				}
			}
			if closerID, ok := t.closedBy[i.ID]; ok {
				pr := &p.prs[p.idMap[closerID]]
				info.ClosedByPR = pr.getCloserInfo()
				if cpIDs, ok := p.cherryPickedTo[pr.ID]; ok {
					for _, cpID := range cpIDs {
						pr := p.prs[p.idMap[cpID]]
						info.CloserCherryPicked = append(info.CloserCherryPicked, pr.getCloserInfo())
					}
				}
			}
			infos = append(infos, info)
		}
	}
	return infos
}
