package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DB struct {
	*sql.DB
}

type Issue struct {
	ID         int
	Owner      string
	Repository string
	Number     int
	Url        string
	Title      string
	Author     string
	Score      string
	Mentor     string
	Hint       string
	Labels     []string
	Assignees  []Assignee
	LinkedPRs  []LinkedPR
}

func (db *DB) GetIssues(state string, labels []string) (result []Issue) {
	args := make([]interface{}, 0, len(labels)+2)
	filter := ""
	args = append(args, state)
	for i, label := range labels {
		if i == 0 {
			filter = "l.name=?"
		} else {
			filter += " or l.name=?"
		}
		args = append(args, label)
	}
	args = append(args, len(labels))
	stmt := fmt.Sprintf("select * from (select i.id, i.owner, i.repository, i.number, i.url, i.title, i.author, i.hint, i.score, i.mentor, count(*) as cnt from issue as i join label as l on i.id = l.issue_id where i.state = ? and (%s) group by i.id order by i.owner, i.repository, i.number desc) s where s.cnt = ?", filter)
	res, err := db.Query(stmt, args...)
	if err != nil {
		log.Fatal(err)
	}
	var i Issue
	var tmp int
	for res.Next() {
		var hint, score, mentor *string
		err := res.Scan(&i.ID, &i.Owner, &i.Repository, &i.Number, &i.Url, &i.Title, &i.Author, &hint, &score, &mentor, &tmp)
		if err != nil {
			log.Fatal(err)
		}
		if hint != nil {
			i.Hint = *hint
		}
		if score != nil {
			i.Score = *score
		}
		if mentor != nil {
			i.Mentor = *mentor
		}
		result = append(result, i)
	}
	res.Close()
	return result
}

func (db *DB) GetIssueLabelsByID(issueID int) []string {
	res, err := db.Query("select name from label where issue_id = ?", issueID)
	if err != nil {
		log.Fatal("error query labels", err)
	}
	result := make([]string, 0)
	var label string
	for res.Next() {
		res.Scan(&label)
		result = append(result, label)
	}
	res.Close()
	return result
}

type Assignee struct {
	Name      string
	CreatedAt time.Time
}

func (db *DB) GetIssueAssigneesByID(issueID int) (result []Assignee) {
	res, err := db.Query("select name, created_at from assignee where issue_id = ?", issueID)
	if err != nil {
		log.Fatal("error query labels", err)
	}
	var ass Assignee
	for res.Next() {
		res.Scan(&ass.Name, &ass.CreatedAt)
		result = append(result, ass)
	}
	res.Close()
	return result
}

type LinkedPR struct {
	ID         int
	Repository string
	Owner      string
	Number     int
	Url        string
	Title      string
	Author     string
}

func (db *DB) GetIssueLinkedPRsByID(issueID int) (result []LinkedPR) {
	res, err := db.Query("select pr.ID, pr.owner, pr.repository, pr.number, pr.url, pr.title, pr.author from pull_request as pr join close as c on c.pull_request_id = pr.id where c.issue_id = ?", issueID)
	if err != nil {
		log.Fatal("error query labels", err)
	}
	var pr LinkedPR
	for res.Next() {
		res.Scan(&pr.ID, &pr.Owner, &pr.Repository, &pr.Number, &pr.Url, &pr.Title, &pr.Author)
		result = append(result, pr)
	}
	res.Close()
	return result
}
