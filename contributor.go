package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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

var pingcapers = []string{"zz-jason", "zhouqiang-cl", "5kbpers", "AndreMouche", "anotherrachel", "AstroProfundis", "BellaXiang", "booooodv", "breeswish", "buggithubs", "CaitinChen", "cfzjywxk", "ChenPeng2013", "Chujie", "CocaLi", "cosven", "csuzhangxc", "cwen0", "cyliu0", "Damon-PingCAP", "dcalvin", "Deardrops", "disksing", "ethercflow", "eurekaka", "francis0407", "g1eny0ung", "GITHUBear", "glorv", "GMHDBJD", "IANTHEREAL", "HunDunDM", "hunterlxt", "husiyu", "iamxy", "innerr", "iosmanthus", "ishiihara", "jackysp", "JaySon-Huang", "kennytm", "kissmydb", "lawyerphx", "leoppro", "lhy1024", "lichunzhu", "lonng", "lysu", "lzmhhh123", "marsishandsome", "meyu44", "Minorli", "NingLin-P", "nolouch", "PiPaSay", "qiuyesuifeng", "qw4990", "ran-huang", "Reminiscent", "sdojjy", "siddontang", "Soline324", "sticnarf", "SunRunAway", "sunzhuohang", "superlzs0476", "sykp241095", "tangenta", "tennix", "tiancaiamao", "together-wang", "toutdesuite", "tshqin", "uglyengineer", "WangXiangUSTC", "wd0517", "Win-Man", "winkyao", "wsabc01", "wshwsh12", "WT-Liu", "xuechunL", "YangKeao", "yikeke", "YiniXu9506", "you06", "youjiali1995", "YuJuncen", "zanmato1984", "zhexuany", "zhongzc", "zyguan", "jebter", "coocood", "imtbkcat", "XuHuaiyu", "fzhedu", "winoros", "crazycs520", "AilinKid", "djshow832", "zimulala", "hanfei1991", "windtalker", "lidezhu", "birdstorm", "solotzg", "MyonKeminta", "nrc", "Connor1996", "brson", "BusyJay", "gengliqi", "overvenus", "hicqu", "Little-Wallace", "yiwu-arbug", "3pointer", "amyangfei", "july2993", "lucklove", "mahjonp", "mapleFU", "shafreeck", "rleungx", "aylei", "DanielZhangQD", "qiffang", "jlerche", "shuijing198799", "weekface", "onlymellb", "LinuxGit", "Yisaer", "cofyc", "baurine", "aytrack", "lilinghai", "ichn-hu", "LittleFall", "leiysky", "andylokandy", "yeya24", "Illyrix", "fewdan", "zhailei9710", "lilin90", "queenypingcap", "WalterWj", "15521174487", "bb7133", "shenli", "dbaoutdo", "huachaohuang", "UncP", "zhangjinpeng1987", "lamxTyler", "liubo0127", "fipped", "wentaoxu", "zhengwanbo", "lhyPingcap", "ciscoxll", "gaohailang", "yanyanqing", "ilovesoup", "dorianzheng", "datahoecn", "UNHNQ", "TomShawn", "xiekeyi98", "chenxiaojing", "ericsyh", "c4pt0r", "pcqz", "Luffbee", "flowbehappy", "ngaut", "hanfei19910905", "sre-bot", "tabokie", "pingcap", "gregwebs", "kolbe", "gingerkidney", "Hoverbear", "Wenting0905", "hashbone", "huangxiuyan", "foreyes", "k-ye", "WenBoYanging", "zyh-hust", "shuke987", "juliezhang1112", "zhenjiaogao", "langyuemeng", "niezefeng", "wowdba", "time-and-fate", "Hexilee", "zjj2wry", "liuzix", "ZenoTan", "ekexium", "nullnotnil", "Rick-lee01", "miaoqingli", "handlerww", "jyz0309", "morgo", "wjhuang2016", "JmPotato", "xuyifangreeneyes", "dyzsr", "xiongjiwei", "bobotu"}

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
		severity  string
		di        float64
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
					severity := "unkown"
					di := 0.0
					for _, l := range labels {
						if strings.HasPrefix(l, "severity/") {
							severity = l[9:]
							break
						}
					}
					switch severity {
					case "critical":
						di = 10
					case "major":
						di = 3
					case "moderate":
						di = 1
					case "minor":
						di = 0.1
					default:
						di = 0
					}
					contrib := Contribution{
						author:    string(pr.Author.Login),
						issue_url: string(n.Url),
						pr_url:    string(pr.Url),
						labels:    labels,
						severity:  severity,
						di:        di,
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
		if lastIssue.Node.UpdatedAt.Time.Before(begin) {
			break
		}
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
	pc := make(map[string]struct{})
	for _, p := range pingcapers {
		pc[p] = struct{}{}
	}
	for _, contri := range list {
		fmt.Println(contri[0].author, len(contri))
		//if _, ok := pc[contri[0].author]; ok {
		//	continue
		//}
		for _, c := range contri {
			_, err := f.WriteString(fmt.Sprintf("%s,%d,%s,%f,%s,%s,%s\n", c.author, len(contri), c.severity, c.di, c.issue_url, c.pr_url, c.closed_at.String()))
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
