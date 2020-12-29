package main

import "github.com/shurcooL/githubv4"

type IssueNode struct {
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
			}
		}
	} `graphql:"timelineItems(first: 50, itemTypes: CROSS_REFERENCED_EVENT)"`
}
