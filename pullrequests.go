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

func getPullRequestsFrom(owner, name string, from time.Time, batchLimit int, totalLimit int) (prs []PullRequest, err error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				Edges []struct {
					Cursor githubv4.String
					Node   PullRequest
				}
			} `graphql:"pullRequests(first: $limit, after: $cursor, orderBy: {field: UPDATED_AT, direction: DESC})"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	cursor := (*githubv4.String)(nil)
	total := 0

	since := from.Add(-1 * time.Minute)
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
		}

		err = client.Query(context.Background(), &query, param)
		if err != nil {
			log.Println(err)
			return
		}
		edges := query.Repository.PullRequests.Edges

		for _, edge := range edges {
			prs = append(prs, edge.Node)
			log.Printf("%06d %s %s\n", edge.Node.Number, edge.Node.UpdatedAt.Format(time.RFC3339), edge.Node.Title)
		}

		cnt := len(edges)
		if cnt != 0 {
			lastIssue := &edges[cnt-1]
			cursor = &lastIssue.Cursor
			lastUpdated := lastIssue.Node.UpdatedAt.Time
			total += cnt
			totalLimit -= cnt
			log.Println(cnt, "fetced", owner, name, lastUpdated)
			if since.After(lastUpdated) {
				break
			}
		}
		if cnt != limit {
			break
		}
	}

	log.Printf("fetched %d pull requests from %s/%s\n", total, owner, name)
	return
}

type TrackedPullRequests struct {
	prs            []PullRequest
	idMap          IDMap
	cherryPickedTo map[githubv4.ID][]githubv4.ID
}

func (t *TrackedPullRequests) Load(data []byte) {
	json.Unmarshal(data, &t.prs)
	log.Printf("load %d prs", len(t.prs))
	t.idMap = make(IDMap)
	for i, pr := range t.prs {
		t.idMap[pr.ID] = i
	}
}

func (ti *TrackedPullRequests) Normalize() {
	sort.Slice(ti.prs, func(i, j int) bool {
		return ti.prs[i].UpdatedAt.Time.After(ti.prs[j].UpdatedAt.Time)
	})
	for i, pr := range ti.prs {
		ti.idMap[pr.ID] = i
	}
}

func (ti *TrackedPullRequests) Save() []byte {
	ti.Normalize()
	issuesJson, err := json.MarshalIndent(ti.prs, "", "\t")
	if err != nil {
		log.Println(err)
	}
	return issuesJson
}

func (t *TrackedPullRequests) add(pr PullRequest) {
	if i, ok := t.idMap[pr.ID]; ok {
		if pr.UpdatedAt.Time.After(t.prs[i].UpdatedAt.Time) {
			t.prs[i] = pr
		}
	} else {
		t.idMap[pr.ID] = len(t.prs)
		t.prs = append(t.prs, pr)
	}
}

func (t *TrackedPullRequests) Add(prs []PullRequest) {
	for _, pr := range prs {
		t.add(pr)
	}
}

func (ti *TrackedPullRequests) getUpdateTimeRange() (from time.Time, to time.Time) {
	if len(ti.prs) == 0 {
		// if we don't have any tracked things, we initialize from 2 days ago
		from = time.Now().Add(-48 * time.Hour)
		to = from
		return
	}
	from = ti.prs[0].UpdatedAt.Time
	to = from
	for _, pr := range ti.prs {
		if pr.UpdatedAt.Time.After(to) {
			to = pr.UpdatedAt.Time
		}
		if from.After(pr.UpdatedAt.Time) {
			from = pr.UpdatedAt.Time
		}
	}
	return
}

func (ti *TrackedPullRequests) Update(from time.Time) (err error, fetched int, added int) {
	var updated []PullRequest
	updated, err = getPullRequestsFrom("pingcap", "tidb", from, 20, 500)
	if err != nil {
		log.Printf("error fetching prs %v", err)
		return
	}
	fetched = len(updated)
	original := len(ti.prs)
	ti.Add(updated)
	added = len(ti.prs) - original
	return
}

func (t *TrackedPullRequests) PopulateCherryPickedTo() {
	t.cherryPickedTo = make(map[githubv4.ID][]githubv4.ID)
	for _, pr := range t.prs {
		for _, edge := range pr.TimelineItems.Edges {
			cpr := edge.Node.CrossReferencedEvent.Source.PullRequest
			if cpr.Number != 0 && strings.HasPrefix(string(cpr.Title), string(pr.Title)) {
				t.cherryPickedTo[pr.ID] = append(t.cherryPickedTo[pr.ID], cpr.ID)
				t.add(PullRequest{PullRequestWithoutTimelineItems: cpr})
			}
		}
	}

	log.Printf("populated %d cherry-picked to relationship", len(t.cherryPickedTo))
}
