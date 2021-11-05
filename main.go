package main

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/olekukonko/tablewriter"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var dbUrl string
var client *githubv4.Client
var trackedIssues map[githubv4.ID]IssueNode
var batchLimit = 100
var debug = false
var report = true
var trackedLabels = [][]string{
	// {"sig/planner"},
	// {"sig/execution"},
	// {"sig/transaction"},
	// {"sig/DDL"},
	{"type/bug"},
}
var PostToIssueID githubv4.ID

func initPostIssueID() error {
	var query struct {
		Repository struct {
			Issue struct {
				ID githubv4.ID
			} `graphql:"issue(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	err := client.Query(context.Background(), &query, map[string]interface{}{
		"name":   githubv4.String("tidb"),
		"owner":  githubv4.String("pingcap"),
		"number": githubv4.Int(20804),
	})
	if err != nil {
		return err
	}
	PostToIssueID = query.Repository.Issue.ID
	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	gt := os.Getenv("GITHUB_TOKEN")
	if gt == "" {
		log.Fatal("no GITHUB_TOKEN found in env")
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gt},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client = githubv4.NewClient(httpClient)
}

// func updateDatabase() {
// 	db, err := sql.Open("mysql", dbUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	re, err := regexp.Compile(`## Score\s*?-\s*?(\d+)\s*?## Mentor\s*?\* @(\w+)`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, issue := range trackedIssues {
// 		// parse score & mentor & hint
// 		// ## Score\n- 300\n\n## Mentor\n * @qw4990
// 		strMatches := re.FindAllStringSubmatch(string(issue.Body), -1)
// 		var score, mentor string
// 		if len(strMatches) != 0 {
// 			idx := len(strMatches) - 1
// 			score, mentor = strMatches[idx][1], strMatches[idx][2]
// 		}
// 		closedAt := &issue.ClosedAt.Time
// 		if closedAt.IsZero() {
// 			closedAt = nil
// 		}
// 		// insert issue
// 		res, err := db.Exec("insert into issue(owner, repository, number, title, author, created_at, updated_at, closed_at, state, url, score, mentor) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) on duplicate key update updated_at=?, closed_at=?, title=?, state=?, score=?, mentor=?",
// 			issue.Repository.Owner.Login, issue.Repository.Name, issue.Number, issue.Title, issue.Author.Login, issue.CreatedAt.Time, issue.UpdatedAt.Time, closedAt, issue.State, issue.Url, score, mentor,
// 			issue.UpdatedAt.Time, closedAt, issue.Title, issue.State, score, mentor)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		issueId, err := res.LastInsertId()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if issueId == 0 {
// 			res, err := db.Query("select id from issue where owner=? and repository=? and number=?", issue.Repository.Owner.Login, issue.Repository.Name, issue.Number)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			res.Next()
// 			if err = res.Scan(&issueId); err != nil {
// 				log.Fatal(err)
// 			}
// 			res.Close()
// 		}

// 		// insert label
// 		if _, err := db.Exec("delete from label where issue_id = ?", issueId); err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, label := range issue.Labels.Nodes {
// 			_, err := db.Exec("insert into label (issue_id, name) values (?, ?) on duplicate key update issue_id=issue_id, name=name", issueId, label.Name)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		// insert assignee
// 		if _, err := db.Exec("delete from assignee where issue_id = ?", issueId); err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, assignee := range issue.Assignees.Nodes {
// 			_, err := db.Exec("insert into assignee (issue_id, name, created_at) values (?, ?, ?) on duplicate key update issue_id=issue_id, name=name", issueId, assignee.Login, assignee.CreatedAt.Time)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		// insert PRs
// 		if _, err := db.Exec("delete from close where issue_id = ?", issueId); err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, edge := range issue.TimelineItems.Edges {
// 			event := edge.Node.CrossReferencedEvent
// 			if !event.WillCloseTarget {
// 				continue
// 			}
// 			pr := event.Source.PullRequest
// 			res, err := db.Exec("insert into pull_request (owner, repository, number, title, author, created_at, updated_at, state, url) values (?, ?, ?, ?, ?, ?, ?, ?, ?) on duplicate key update updated_at=?, title=?, state=?", pr.Repository.Owner.Login, pr.Repository.Name, pr.Number, pr.Title, pr.Author.Login, pr.CreatedAt.Time, pr.UpdatedAt.Time, pr.State, pr.Url, pr.UpdatedAt.Time, pr.Title, pr.State)
// 			if err != nil {
// 				log.Fatal("failed on insert PR", err)
// 			}
// 			prId, err := res.LastInsertId()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			if prId == 0 {
// 				res, err := db.Query("select id from pull_request where owner=? and repository=? and number=?", pr.Repository.Owner.Login, pr.Repository.Name, pr.Number)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				res.Next()
// 				if err = res.Scan(&prId); err != nil {
// 					log.Fatal(err)
// 				}
// 				res.Close()
// 			}
// 			if _, err := db.Exec("insert into close(issue_id, pull_request_id) values (?, ?)", issueId, prId); err != nil {
// 				log.Fatal("failed insert close", err)
// 			}
// 		}

// 	}
// }

func generateTrackTable() string {
	mysqlDB, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db := DB{mysqlDB}

	tables := make([]string, 0)

	for _, labels := range trackedLabels {
		issues := db.GetIssues("OPEN", labels)
		header := []string{"issue", "priority", "assignee", "pr", "hint"}
		data := make([][]string, 0, len(issues))

		for _, i := range issues {
			i.Labels = db.GetIssueLabelsByID(i.ID)
			i.Assignees = db.GetIssueAssigneesByID(i.ID)
			i.LinkedPRs = db.GetIssueLinkedPRsByID(i.ID)
			isBug := false
			for _, l := range i.Labels {
				if l == "type/bug" {
					isBug = true
				}
			}
			if !isBug {
				continue
			}
			picked := false
			assigned := false
			linked := false
			d := make([]string, 0, len(header))
			// issue
			d = append(d, fmt.Sprintf("[#%d](%s)", i.Number, i.Url))

			// challenge
			d = append(d, func(i Issue) string {
				helpWanted := false
				for _, l := range i.Labels {
					if l == "challenge-program" {
						helpWanted = true
					}
					if l == "picked" {
						picked = true
					}
				}
				var content string
				if helpWanted {
					if picked || len(i.Assignees) != 0 {
						content = "&#x2B50; picked"
					} else {
						content = "&#x2665; yes!"
					}
					if len(i.Mentor) != 0 {
						content += "</br>Mentor: @" + i.Mentor
					}
					if len(i.Score) != 0 {
						content += "</br>Score: " + i.Score
					}
					return content
				}
				return ""
			}(i))
			// remove challenge
			d = d[:len(d)-1]

			// priority
			d = append(d, func(labels []string) string {
				for _, l := range labels {
					if l == "severity/critical" {
						return "critical"
					}
					if l == "severity/major" {
						return "major"
					}
					if l == "severity/moderate" {
						return "moderate"
					}
					if l == "severity/minor" {
						return "minor"
					}
				}
				return ""
			}(i.Labels))

			// assignee
			d = append(d, strings.Join(func(a []Assignee) (result []string) {
				for i := range a {
					name := "@" + a[i].Name
					if len(a[i].Name) > 9 {
						name = "<sub>@" + a[i].Name + "</sub>"
					}
					if len(a[i].Name) > 12 {
						name = "<sub><sup>@" + a[i].Name + "</sup></sub>"
					}
					result = append(result, name)
					assigned = true
				}
				return
			}(i.Assignees), "</br>"))

			// pr
			d = append(d, strings.Join(func(a []LinkedPR) (result []string) {
				for i := range a {
					result = append(result, fmt.Sprintf("[#%d](%s)", a[i].Number, a[i].Url))
					linked = true
				}
				return
			}(i.LinkedPRs), "</br>"))

			// hint
			d = append(d, i.Hint)
			if !assigned && !picked && !linked {
				d[0] = d[0] + "&#x2757;"
			}
			data = append(data, d)
		}
		order := map[string]int{
			"critical": 1,
			"major":    2,
			"moderate": 3,
			"minor":    4,
		}
		sort.SliceStable(data, func(i, j int) bool {
			ii, iok := order[data[i][2]]
			jj, jok := order[data[j][2]]
			if !iok && !jok {
				return true
			}
			if !iok {
				return false
			}
			if !jok {
				return true
			}
			return ii < jj
		})
		var buf bytes.Buffer
		table := tablewriter.NewWriter(&buf)
		table.SetHeader(header)
		table.SetColWidth(100000) // don't break line
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data)
		table.Render()
		tables = append(tables, string(buf.Bytes()))

	}

	type Tables struct {
		TableSigPlanner     string
		TableSigExecution   string
		TableSigTransaction string
		TableSigDDL         string
		UpdatedAt           string
	}
	tp := `

If you are interested in database development, or you are a TiDB user, no matter what, if you want to contribute to TiDB and learn about how a distributed HTAP database worked, here is the right place.

Let's get started by solving some bugs! Here is a curated list of some easy-to-go bugs, pick the one that you want to smash!

* [sig/planner](#sig/planner)
* [sig/execution](#sig/execution)
* [sig/transaction](#sig/transaction)
* [sig/DDL](#sig/DDL)

Note: currently the issues are classified by their SIG owners, such as sig/planner and sig/execution, which stands for special interests groups that focus on SQL planning and SQL execution, to know more about TiDB community, see [the community repository](https://github.com/pingcap/community). We also host discussions on slack, if you are not in the corresponding slack channel, we highly recommend you to join so that you could ask questions and get responses immediately from these SIGs members. [Join TiDB Community slack workspace now!](https://join.slack.com/t/tidbcommunity/shared_invite/enQtNzc0MzI4ODExMDc4LWYwYmIzMjZkYzJiNDUxMmZlN2FiMGJkZjAyMzQ5NGU0NGY0NzI3NTYwMjAyNGQ1N2I2ZjAxNzc1OGUwYWM0NzE)

If there is a &#x2757; before the issue link, it means there is no one assigned, nor a PR linked, nor picked, and it is for the maintainers to track the progress of each issue, it is also a notation of "welcome to take a look".

Feel free to comment on issues that interest you, and ask whatever questions you have on how to get started working on them!

<h2 name="sig/planner">sig/planner</h2>

{{ .TableSigPlanner }}

<h2 name="sig/execution">sig/execution</h2>

{{ .TableSigExecution }}

<h2 name="sig/transaction">sig/transaction</h2>

{{ .TableSigTransaction }}

<h2 name="sig/DDL">sig/DDL</h2>

{{ .TableSigDDL }}


---

updated at {{ .UpdatedAt }}

`
	now := time.Now()
	tbs := Tables{tables[0], tables[1], tables[2], tables[3], fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())}
	t, err := template.New("content").Parse(tp)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, tbs)
	content := strings.Join(tables, "\n\n")
	content = string(buf.Bytes())
	log.Println(content)
	if err := ioutil.WriteFile("index.md", []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
	return content
}

func reportToIssue(content string) (url string, err error) {
	title := "Welcome to contribute"

	var m struct {
		UpdateIssue struct {
			Issue struct {
				Url githubv4.String
			}
		} `graphql:"updateIssue(input: $input)"`
	}
	input := githubv4.UpdateIssueInput{
		ID:    PostToIssueID,
		Title: githubv4.NewString(githubv4.String(title)),
		Body:  githubv4.NewString(githubv4.String(content)),
	}
	err = client.Mutate(context.Background(), &m, input, nil)
	if err != nil {
		return
	}
	url = string(m.UpdateIssue.Issue.Url)
	return
}

func update(ti *TrackedIssues, tpr *TrackedPullRequests, extend int) {
	tiFrom, tiTo := ti.getUpdateTimeRange()

	now := time.Now()
	err, fetched, added := ti.UpdateByTimeRange(tiTo, now)
	if err != nil {
		log.Println("failed to update issue", err)
		return
	}
	log.Printf("fetched %d added %d to tracked issues", fetched, added)
	err, fetched, added = tpr.Update(tiTo)
	if err != nil {
		log.Println("failed to update pr", err)
		return
	}
	log.Printf("fetched %d added %d to tracked prs", fetched, added)

	accumulated := 0
	extendBy := 48 * time.Hour
	for accumulated < extend {
		earlier := tiFrom.Add(-extendBy)
		err, fetched, added = ti.UpdateByTimeRange(earlier, tiFrom)
		tiFrom = earlier
		if err != nil {
			log.Println("failed to update")
			break
		}
		accumulated += fetched
		log.Printf("fetched %d added %d to tracked issue\n", fetched, added)
		log.Printf("issue tracked from %s\n", tiFrom)
		log.Printf("accumulated fetching %d limit by %d\n", accumulated, extend)
	}
}

func readFileFromZip(fp string) (map[string][]byte, error) {
	r, err := zip.OpenReader(fp)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	files := make(map[string][]byte)

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		files[f.Name], err = ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		rc.Close()
	}
	return files, nil
}

func writeFileToZip(fp string, files map[string][]byte) error {
	zipFile, err := os.Create(fp)
	if err != nil {
		return err
	}
	zipFileWriter := zip.NewWriter(zipFile)
	defer zipFileWriter.Close()

	for name, content := range files {
		fileWriter, err := zipFileWriter.Create(name)
		if err != nil {
			return err
		}
		contentLen := len(content)
		for contentLen > 0 {
			n, err := fileWriter.Write(content)
			if err != nil {
				return err
			}
			contentLen -= n
			content = content[n:]
		}
	}
	return nil
}

func main() {
	getContri := flag.Bool("contri", false, "get contributors")
	getIssueInfo := flag.Int("issue", 0, "the number of the issue to be examined")
	runUpdate := flag.Bool("update", false, "if run update")
	numExtend := flag.Int("extend", 0, "the number of issues to extend back in history")
	flag.Parse()

	archiveFilePath := "raw.zip"
	issuesPath := "issues.json"
	prsPath := "prs.json"

	ti := &TrackedIssues{}
	tpr := &TrackedPullRequests{}
	start := time.Now()

	fileData, err := readFileFromZip(archiveFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if issuesData, ok := fileData[issuesPath]; ok {
		ti.Load(issuesData)
	} else {
		log.Println("no issues data")
	}

	if prsData, ok := fileData[prsPath]; ok {
		tpr.Load(prsData)
	} else {
		log.Println("no prs data")
	}

	log.Printf("data loaded in %v", time.Now().Sub(start))

	if *runUpdate {
		update(ti, tpr, *numExtend)

		files := make(map[string][]byte)
		files[issuesPath] = ti.Save()
		files[prsPath] = tpr.Save()
		tmpFilePath := "tmp.zip"
		if err := writeFileToZip(tmpFilePath, files); err != nil {
			log.Println(err)
		} else {
			// atomically replace the old zip file
			if err := os.Rename(tmpFilePath, archiveFilePath); err != nil {
				log.Println(err)
			}
		}
	}

	ti.PopulateClosedBy(tpr)
	tpr.PopulateCherryPickedTo()
	log.Printf("%d issues and %d prs in track", len(ti.issues), len(tpr.prs))

	infos := GetClosedIssueInfo(ti, tpr)
	data, err := json.MarshalIndent(infos, "", "\t")
	if err != nil {
		log.Println(err)
	}
	ioutil.WriteFile("infos.json", data, 0644)

	if *getContri {
		err := getContributors()
		if err != nil {
			panic(err)
		}
	} else if *getIssueInfo != 0 {
		fmt.Printf("okay to track")
	} else {
		// initForTrack()
		// fmt.Println(len(trackedIssues))
		// updateDatabase()
		// content := generateTrackTable()
		// _ = content
		// if report {
		// 	reportToIssue(content)
		// }
	}
}
