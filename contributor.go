package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/shurcooL/githubv4"
)

type IssueNode2 struct {
	Title  githubv4.String
	State  githubv4.IssueState
	ID     githubv4.ID
	Number githubv4.Int
	Url    githubv4.String
	Author struct {
		Login githubv4.String
	}
	Body       githubv4.String
	ClosedAt   githubv4.DateTime
	CreatedAt  githubv4.DateTime
	UpdatedAt  githubv4.DateTime
	Repository struct {
		Name  githubv4.String
		Owner struct {
			Login githubv4.String
		}
	}
	Labels struct {
		Nodes []struct {
			Name githubv4.String
		}
	} `graphql:"labels(last: 15)"`
	Assignees struct {
		Nodes []struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	} `graphql:"assignees(last: 15)"`
	TimelineItems struct {
		Edges []struct {
			Node struct {
				CrossReferencedEvent struct {
					WillCloseTarget githubv4.Boolean
					Source          struct {
						PullRequest struct {
							State  githubv4.PullRequestState
							Author struct {
								Login githubv4.String
							}
							CreatedAt githubv4.DateTime
							UpdatedAt githubv4.DateTime
							Title     githubv4.String
							Url       githubv4.String
							Number    githubv4.Int
							Labels    struct {
								Nodes []struct {
									Name githubv4.String
								}
							} `graphql:"labels(last: 15)"`
							Repository struct {
								Name  githubv4.String
								Owner struct {
									Login githubv4.String
								}
							}
						} `graphql:"... on PullRequest"`
					}
				} `graphql:"... on CrossReferencedEvent"`
				ClosedEvent struct {
					CreatedAt githubv4.DateTime
					Closer    struct {
						PullRequest struct {
							State  githubv4.PullRequestState
							Author struct {
								Login githubv4.String
							}
							CreatedAt githubv4.DateTime
							UpdatedAt githubv4.DateTime
							Title     githubv4.String
							Url       githubv4.String
							Number    githubv4.Int
							Labels    struct {
								Nodes []struct {
									Name githubv4.String
								}
							} `graphql:"labels(last: 15)"`
							Repository struct {
								Name  githubv4.String
								Owner struct {
									Login githubv4.String
								}
							}
						} `graphql:"... on PullRequest"`
					}
				} `graphql:"... on ClosedEvent"`
			}
		}
	} `graphql:"timelineItems(first: 50, itemTypes: CLOSED_EVENT)"`
}

func getContributors() error {
	owner := "pingcap"
	name := "tidb"
	var query struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Cursor githubv4.String
					Node   IssueNode2
				}
			} `graphql:"issues(first: $limit, after: $cursor, states: CLOSED, orderBy: {field: UPDATED_AT, direction: DESC}, labels: $labels)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}
	cursor := (*githubv4.String)(nil)
	total := 0
	labels := []string{"type/bug"}
	ghLabels := make([]githubv4.String, 0, len(labels))
	for _, l := range labels {
		ghLabels = append(ghLabels, githubv4.String(l))
	}
	type Contribution struct {
		author    string
		issue_url string
		pr_url    string
		labels    []string
		closed_at time.Time
	}
	contributions := make(map[string][]Contribution)
	begin := time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)
	for {
		param := map[string]interface{}{
			"name":   githubv4.String(name),
			"owner":  githubv4.String(owner),
			"limit":  githubv4.Int(batchLimit),
			"cursor": cursor,
			"labels": ghLabels,
		}
		err := client.Query(context.Background(), &query, param)
		if err != nil {
			log.Println(err)
			return err
		}
		issues := query.Repository.Issues.Edges
		cnt := len(issues)
		lastIssue := &issues[cnt-1]
		if cnt != 0 {
			cursor = &lastIssue.Cursor
			//lastUpdated := lastIssue.Node.UpdatedAt
			//if latestUpdate.Sub(lastUpdated.Time) > 2*time.Hour {
			//	break
			//}
		}
		total += cnt
		log.Println(cnt, "fetced", owner, name, labels)

		for _, issue := range issues {
			n := issue.Node
			for _, event := range n.TimelineItems.Edges {
				pr := event.Node.ClosedEvent.Closer.PullRequest
				if string(pr.State) == "MERGED" {
					labels := make([]string, len(n.Labels.Nodes))
					for i, l := range n.Labels.Nodes {
						labels[i] = string(l.Name)
					}
					contrib := Contribution{
						author:    string(pr.Author.Login),
						issue_url: string(n.Url),
						pr_url:    string(pr.Url),
						labels:    labels,
						closed_at: event.Node.ClosedEvent.CreatedAt.Time,
					}
					if contrib.closed_at.After(end) || contrib.closed_at.Before(begin) {
						continue
					}
					c, has := contributions[string(pr.Author.Login)]
					if !has {
						c = []Contribution{contrib}
					} else {
						c = append(c, contrib)
					}
					contributions[string(pr.Author.Login)] = c
				}
			}
		}

		if cnt != batchLimit {
			break
		}
		if debug {
			break
		}
		//		if lastIssue.Node.UpdatedAt.Time.Before(begin) {
		//			break
		//		}
	}
	fmt.Println("hello contributor")
	list := make([][]Contribution, 0, len(contributions))
	for _, contri := range contributions {
		list = append(list, contri)
	}
	sort.Slice(list, func(i, j int) bool {
		return len(list[i]) > len(list[j])
	})
	f, err := os.Create("./contributor.csv")
	if err != nil {
		panic(err)
	}
	for _, contri := range list {
		fmt.Println(contri[0].author, len(contri))
		for _, c := range contri {
			_, err := f.WriteString(fmt.Sprintf("%s,%d,%s,%s,%s\n", c.author, len(contri), c.issue_url, c.pr_url, c.closed_at.String()))
			if err != nil {
				return err
			}
		}
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
