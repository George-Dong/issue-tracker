

If you are interested in database development, or you are a TiDB user, no matter what, if you want to contribute to TiDB and learn about how a distributed HTAP database worked, here is the right place.

Let's get started by solving some bugs! Here is a curated list of some easy-to-go bugs, pick the one that you want to smash!

* [sig/planner](#sig/planner)
* [sig/execution](#sig/execution)
* [sig/transaction](#sig/transaction)
* [sig/DDL](#sig/DDL)

Note: currently the issues are classified by their SIG owners, such as sig/planner and sig/execution which stands for special interests groups that focus on SQL planning and SQL execution, to know more about TiDB community, see [the community repository](https://github.com/pingcap/community). We also host discussions on slack, if you are not in the corresponding slack channel, we highly recommend you to join so that you could ask questions and get responses immediately from these SIGs members. [Join TiDB Community slack workspace now!](https://join.slack.com/t/tidbcommunity/shared_invite/enQtNzc0MzI4ODExMDc4LWYwYmIzMjZkYzJiNDUxMmZlN2FiMGJkZjAyMzQ5NGU0NGY0NzI3NTYwMjAyNGQ1N2I2ZjAxNzc1OGUwYWM0NzE)

The "challenge" column in the following tables indicates whether the issue is of a challenge program, and "yes" means it is, while "picked" means it has been picked by someone.

If there is a &#x2757; before the issue link, it means there is no one assigned, nor a PR linked, nor picked, and it is for the maintainers to track the progress of each issue, it is also a notation of "welcome to take a look".

<h2 name="sig/planner">sig/planner</h2>

|                             ISSUE                              |                       CHALLENGE                        | PRIORITY |                ASSIGNEE                 |                          PR                          | HINT |
|----------------------------------------------------------------|--------------------------------------------------------|----------|-----------------------------------------|------------------------------------------------------|------|
| [#22404](https://github.com/pingcap/tidb/issues/22404)&#x2757; |                                                        | major    |                                         |                                                      |      |
| [#22396](https://github.com/pingcap/tidb/issues/22396)         |                                                        | major    | <sub><sup>@beihaiguaishou</sup></sub>   | [#22452](https://github.com/pingcap/tidb/pull/22452) |      |
| [#22388](https://github.com/pingcap/tidb/issues/22388)         |                                                        | major    | @rebelice                               |                                                      |      |
| [#22301](https://github.com/pingcap/tidb/issues/22301)         |                                                        | major    | <sub><sup>@xuyifangreeneyes</sup></sub> |                                                      |      |
| [#21677](https://github.com/pingcap/tidb/issues/21677)         |                                                        | major    | <sub><sup>@xuyifangreeneyes</sup></sub> |                                                      |      |
| [#21035](https://github.com/pingcap/tidb/issues/21035)         |                                                        | major    | @qw4990                                 |                                                      |      |
| [#22082](https://github.com/pingcap/tidb/issues/22082)&#x2757; |                                                        | moderate |                                         |                                                      |      |
| [#21475](https://github.com/pingcap/tidb/issues/21475)&#x2757; |                                                        | moderate |                                         |                                                      |      |
| [#21098](https://github.com/pingcap/tidb/issues/21098)&#x2757; | &#x2665; yes!                                          | moderate |                                         |                                                      |      |
| [#20893](https://github.com/pingcap/tidb/issues/20893)         |                                                        | moderate | <sub>@iosmanthus</sub></br>@wshwsh12    | [#21444](https://github.com/pingcap/tidb/pull/21444) |      |
| [#19231](https://github.com/pingcap/tidb/issues/19231)&#x2757; |                                                        | moderate |                                         |                                                      |      |
| [#18216](https://github.com/pingcap/tidb/issues/18216)         |                                                        | moderate | @fzhedu                                 |                                                      |      |
| [#18042](https://github.com/pingcap/tidb/issues/18042)         | &#x2665; yes!</br>Mentor: @wshwsh12</br>Score: 300     | moderate |                                         | [#20905](https://github.com/pingcap/tidb/pull/20905) |      |
| [#17852](https://github.com/pingcap/tidb/issues/17852)         |                                                        | moderate | @XuHuaiyu                               |                                                      |      |
| [#17731](https://github.com/pingcap/tidb/issues/17731)         | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | moderate |                                         | [#22416](https://github.com/pingcap/tidb/pull/22416) |      |
| [#16909](https://github.com/pingcap/tidb/issues/16909)         |                                                        | moderate | @lzmhhh123                              | [#17046](https://github.com/pingcap/tidb/pull/17046) |      |
| [#16788](https://github.com/pingcap/tidb/issues/16788)         | &#x2B50; picked</br>Mentor: @SunRunAway</br>Score: 300 | moderate | <sub><sup>@xuyifangreeneyes</sup></sub> | [#20044](https://github.com/pingcap/tidb/pull/20044) |      |
| [#16764](https://github.com/pingcap/tidb/issues/16764)         |                                                        | moderate | @lzmhhh123                              |                                                      |      |
| [#16062](https://github.com/pingcap/tidb/issues/16062)         | &#x2665; yes!</br>Mentor: @qw4990</br>Score: 300       | moderate |                                         | [#18218](https://github.com/pingcap/tidb/pull/18218) |      |
| [#13856](https://github.com/pingcap/tidb/issues/13856)         |                                                        | moderate | @winoros                                |                                                      |      |
| [#13167](https://github.com/pingcap/tidb/issues/13167)         |                                                        | moderate | @fzhedu                                 |                                                      |      |
| [#12182](https://github.com/pingcap/tidb/issues/12182)&#x2757; | &#x2665; yes!                                          | moderate |                                         |                                                      |      |
| [#10151](https://github.com/pingcap/tidb/issues/10151)         | &#x2B50; picked</br>Mentor: @lzmhhh123</br>Score: 300  | moderate | <sub>@wjhuang2016</sub>                 |                                                      |      |
| [#9373](https://github.com/pingcap/tidb/issues/9373)           |                                                        | moderate | @morgo                                  |                                                      |      |
| [#8190](https://github.com/pingcap/tidb/issues/8190)           | &#x2B50; picked                                        | moderate | <sub><sup>@xuyifangreeneyes</sup></sub> |                                                      |      |
| [#22535](https://github.com/pingcap/tidb/issues/22535)&#x2757; |                                                        | minor    |                                         |                                                      |      |
| [#21625](https://github.com/pingcap/tidb/issues/21625)&#x2757; |                                                        | minor    |                                         |                                                      |      |
| [#21454](https://github.com/pingcap/tidb/issues/21454)&#x2757; |                                                        | minor    |                                         |                                                      |      |
| [#20019](https://github.com/pingcap/tidb/issues/20019)&#x2757; |                                                        | minor    |                                         |                                                      |      |
| [#19802](https://github.com/pingcap/tidb/issues/19802)&#x2757; | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                                         |                                                      |      |
| [#19713](https://github.com/pingcap/tidb/issues/19713)         |                                                        | minor    | @winoros                                |                                                      |      |
| [#16473](https://github.com/pingcap/tidb/issues/16473)         |                                                        | minor    | @winoros                                |                                                      |      |
| [#15514](https://github.com/pingcap/tidb/issues/15514)&#x2757; | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                                         |                                                      |      |
| [#11977](https://github.com/pingcap/tidb/issues/11977)&#x2757; | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                                         |                                                      |      |


<h2 name="sig/execution">sig/execution</h2>

|                             ISSUE                              |                       CHALLENGE                        | PRIORITY |        ASSIGNEE         |                                                      PR                                                       | HINT |
|----------------------------------------------------------------|--------------------------------------------------------|----------|-------------------------|---------------------------------------------------------------------------------------------------------------|------|
| [#22530](https://github.com/pingcap/tidb/issues/22530)&#x2757; |                                                        | major    |                         |                                                                                                               |      |
| [#22525](https://github.com/pingcap/tidb/issues/22525)&#x2757; |                                                        | major    |                         |                                                                                                               |      |
| [#22496](https://github.com/pingcap/tidb/issues/22496)         |                                                        | major    | @jyz0309                | [#22507](https://github.com/pingcap/tidb/pull/22507)                                                          |      |
| [#22398](https://github.com/pingcap/tidb/issues/22398)&#x2757; |                                                        | major    |                         |                                                                                                               |      |
| [#21584](https://github.com/pingcap/tidb/issues/21584)&#x2757; |                                                        | major    |                         |                                                                                                               |      |
| [#19025](https://github.com/pingcap/tidb/issues/19025)         |                                                        | major    | <sub>@SunRunAway</sub>  | [#19029](https://github.com/pingcap/tidb/pull/19029)                                                          |      |
| [#22447](https://github.com/pingcap/tidb/issues/22447)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22446](https://github.com/pingcap/tidb/issues/22446)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22445](https://github.com/pingcap/tidb/issues/22445)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22444](https://github.com/pingcap/tidb/issues/22444)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22442](https://github.com/pingcap/tidb/issues/22442)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22394](https://github.com/pingcap/tidb/issues/22394)         |                                                        | moderate |                         | [#22407](https://github.com/pingcap/tidb/pull/22407)                                                          |      |
| [#22390](https://github.com/pingcap/tidb/issues/22390)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22389](https://github.com/pingcap/tidb/issues/22389)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22387](https://github.com/pingcap/tidb/issues/22387)         |                                                        | moderate | @Tjianke                | [#22426](https://github.com/pingcap/tidb/pull/22426)                                                          |      |
| [#22386](https://github.com/pingcap/tidb/issues/22386)&#x2757; |                                                        | moderate |                         |                                                                                                               |      |
| [#22206](https://github.com/pingcap/tidb/issues/22206)&#x2757; |                                                        | moderate |                         |                                                                                                               |      |
| [#22096](https://github.com/pingcap/tidb/issues/22096)&#x2757; |                                                        | moderate |                         |                                                                                                               |      |
| [#22088](https://github.com/pingcap/tidb/issues/22088)&#x2757; |                                                        | moderate |                         |                                                                                                               |      |
| [#21653](https://github.com/pingcap/tidb/issues/21653)         |                                                        | moderate |                         | [#21230](https://github.com/pingcap/tidb/pull/21230)                                                          |      |
| [#20563](https://github.com/pingcap/tidb/issues/20563)         | &#x2B50; picked                                        | moderate | <sub>@wjhuang2016</sub> | [#1129](https://github.com/pingcap/parser/pull/1129)                                                          |      |
| [#20411](https://github.com/pingcap/tidb/issues/20411)         |                                                        | moderate | @qw4990                 |                                                                                                               |      |
| [#18488](https://github.com/pingcap/tidb/issues/18488)         |                                                        | moderate | @morgo                  | [#19029](https://github.com/pingcap/tidb/pull/19029)                                                          |      |
| [#18309](https://github.com/pingcap/tidb/issues/18309)         |                                                        | moderate | @lzmhhh123              |                                                                                                               |      |
| [#17677](https://github.com/pingcap/tidb/issues/17677)         |                                                        | moderate | @qw4990                 |                                                                                                               |      |
| [#17486](https://github.com/pingcap/tidb/issues/17486)         | &#x2B50; picked</br>Mentor: @lzmhhh123</br>Score: 300  | moderate | @morgo                  |                                                                                                               |      |
| [#17083](https://github.com/pingcap/tidb/issues/17083)         |                                                        | moderate | @qw4990                 |                                                                                                               |      |
| [#16679](https://github.com/pingcap/tidb/issues/16679)         | &#x2B50; picked</br>Mentor: @XuHuaiyu</br>Score: 300   | moderate |                         | [#20040](https://github.com/pingcap/tidb/pull/20040)                                                          |      |
| [#15884](https://github.com/pingcap/tidb/issues/15884)&#x2757; | &#x2665; yes!</br>Mentor: @qw4990</br>Score: 300       | moderate |                         |                                                                                                               |      |
| [#15608](https://github.com/pingcap/tidb/issues/15608)         |                                                        | moderate | @ichn-hu                |                                                                                                               |      |
| [#15234](https://github.com/pingcap/tidb/issues/15234)         | &#x2B50; picked</br>Mentor: @XuHuaiyu</br>Score: 300   | moderate |                         | [#20015](https://github.com/pingcap/tidb/pull/20015)                                                          |      |
| [#14399](https://github.com/pingcap/tidb/issues/14399)&#x2757; |                                                        | moderate |                         |                                                                                                               |      |
| [#13157](https://github.com/pingcap/tidb/issues/13157)         | &#x2B50; picked</br>Mentor: @XuHuaiyu</br>Score: 300   | moderate | <sub>@wjhuang2016</sub> |                                                                                                               |      |
| [#11932](https://github.com/pingcap/tidb/issues/11932)         |                                                        | moderate | @wshwsh12               |                                                                                                               |      |
| [#11866](https://github.com/pingcap/tidb/issues/11866)         | &#x2B50; picked</br>Mentor: @qw4990</br>Score: 300     | moderate | @dragonly               |                                                                                                               |      |
| [#22477](https://github.com/pingcap/tidb/issues/22477)&#x2757; |                                                        | minor    |                         |                                                                                                               |      |
| [#22423](https://github.com/pingcap/tidb/issues/22423)         |                                                        | minor    |                         | [#22432](https://github.com/pingcap/tidb/pull/22432)                                                          |      |
| [#22399](https://github.com/pingcap/tidb/issues/22399)&#x2757; |                                                        | minor    |                         |                                                                                                               |      |
| [#22227](https://github.com/pingcap/tidb/issues/22227)&#x2757; |                                                        | minor    |                         |                                                                                                               |      |
| [#21787](https://github.com/pingcap/tidb/issues/21787)&#x2757; |                                                        | minor    |                         |                                                                                                               |      |
| [#21374](https://github.com/pingcap/tidb/issues/21374)         | &#x2665; yes!                                          | minor    |                         | [#21508](https://github.com/pingcap/tidb/pull/21508)                                                          |      |
| [#21307](https://github.com/pingcap/tidb/issues/21307)&#x2757; |                                                        | minor    |                         |                                                                                                               |      |
| [#20018](https://github.com/pingcap/tidb/issues/20018)&#x2757; | &#x2665; yes!</br>Mentor: @SunRunAway</br>Score: 300   | minor    |                         |                                                                                                               |      |
| [#18493](https://github.com/pingcap/tidb/issues/18493)         |                                                        | minor    | @fzhedu                 |                                                                                                               |      |
| [#18401](https://github.com/pingcap/tidb/issues/18401)         |                                                        | minor    | @hanlins                | [#21340](https://github.com/pingcap/tidb/pull/21340)                                                          |      |
| [#17993](https://github.com/pingcap/tidb/issues/17993)&#x2757; | &#x2665; yes!</br>Mentor: @qw4990</br>Score: 300       | minor    |                         |                                                                                                               |      |
| [#17832](https://github.com/pingcap/tidb/issues/17832)         | &#x2B50; picked</br>Mentor: @SunRunAway</br>Score: 300 | minor    | @qw4990                 | [#19854](https://github.com/pingcap/tidb/pull/19854)</br>[#22432](https://github.com/pingcap/tidb/pull/22432) |      |
| [#17751](https://github.com/pingcap/tidb/issues/17751)&#x2757; | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                         |                                                                                                               |      |
| [#17489](https://github.com/pingcap/tidb/issues/17489)         | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                         | [#22475](https://github.com/pingcap/tidb/pull/22475)                                                          |      |
| [#17008](https://github.com/pingcap/tidb/issues/17008)&#x2757; | &#x2665; yes!</br>Mentor: @lzmhhh123</br>Score: 300    | minor    |                         |                                                                                                               |      |
| [#13440](https://github.com/pingcap/tidb/issues/13440)         |                                                        | minor    | <sub>@SunRunAway</sub>  |                                                                                                               |      |
| [#13018](https://github.com/pingcap/tidb/issues/13018)         |                                                        | minor    | @qw4990                 |                                                                                                               |      |
| [#11200](https://github.com/pingcap/tidb/issues/11200)&#x2757; | &#x2665; yes!</br>Mentor: @qw4990</br>Score: 300       | minor    |                         |                                                                                                               |      |
| [#10543](https://github.com/pingcap/tidb/issues/10543)&#x2757; | &#x2665; yes!                                          | minor    |                         |                                                                                                               |      |
| [#8205](https://github.com/pingcap/tidb/issues/8205)&#x2757;   | &#x2665; yes!                                          | minor    |                         |                                                                                                               |      |


<h2 name="sig/transaction">sig/transaction</h2>

|                             ISSUE                              | CHALLENGE | PRIORITY |         ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|-----------|----------|--------------------------|------------------------------------------------------|------|
| [#22400](https://github.com/pingcap/tidb/issues/22400)         |           | major    | <sub>@youjiali1995</sub> | [#22449](https://github.com/pingcap/tidb/pull/22449) |      |
| [#22345](https://github.com/pingcap/tidb/issues/22345)         |           | major    | @lysu                    | [#22372](https://github.com/pingcap/tidb/pull/22372) |      |
| [#21618](https://github.com/pingcap/tidb/issues/21618)         |           | major    |                          | [#21641](https://github.com/pingcap/tidb/pull/21641) |      |
| [#21509](https://github.com/pingcap/tidb/issues/21509)         |           | major    |                          | [#21641](https://github.com/pingcap/tidb/pull/21641) |      |
| [#21335](https://github.com/pingcap/tidb/issues/21335)         |           | major    | @you06                   | [#22146](https://github.com/pingcap/tidb/pull/22146) |      |
| [#21087](https://github.com/pingcap/tidb/issues/21087)         |           | major    | @you06                   |                                                      |      |
| [#20028](https://github.com/pingcap/tidb/issues/20028)         |           | major    | <sub>@tiancaiamao</sub>  | [#21148](https://github.com/pingcap/tidb/pull/21148) |      |
| [#10657](https://github.com/pingcap/tidb/issues/10657)         |           | major    | @nolouch                 |                                                      |      |
| [#9762](https://github.com/pingcap/tidb/issues/9762)&#x2757;   |           | major    |                          |                                                      |      |
| [#22395](https://github.com/pingcap/tidb/issues/22395)         |           | moderate | @ekexium                 | [#22500](https://github.com/pingcap/tidb/pull/22500) |      |
| [#22393](https://github.com/pingcap/tidb/issues/22393)&#x2757; |           | moderate |                          |                                                      |      |
| [#22356](https://github.com/pingcap/tidb/issues/22356)&#x2757; |           | moderate |                          |                                                      |      |
| [#22244](https://github.com/pingcap/tidb/issues/22244)&#x2757; |           | moderate |                          |                                                      |      |
| [#21688](https://github.com/pingcap/tidb/issues/21688)         |           | moderate | @you06                   | [#21878](https://github.com/pingcap/tidb/pull/21878) |      |
| [#21506](https://github.com/pingcap/tidb/issues/21506)         |           | moderate | @cfzjywxk                |                                                      |      |
| [#21355](https://github.com/pingcap/tidb/issues/21355)&#x2757; |           | moderate |                          |                                                      |      |
| [#20990](https://github.com/pingcap/tidb/issues/20990)&#x2757; |           | moderate |                          |                                                      |      |
| [#18048](https://github.com/pingcap/tidb/issues/18048)         |           | moderate | @qw4990                  |                                                      |      |
| [#17798](https://github.com/pingcap/tidb/issues/17798)&#x2757; |           | moderate |                          |                                                      |      |
| [#17797](https://github.com/pingcap/tidb/issues/17797)&#x2757; |           | moderate |                          |                                                      |      |
| [#13958](https://github.com/pingcap/tidb/issues/13958)         |           | moderate | @fzhedu                  |                                                      |      |
| [#22516](https://github.com/pingcap/tidb/issues/22516)&#x2757; |           | minor    |                          |                                                      |      |
| [#20949](https://github.com/pingcap/tidb/issues/20949)&#x2757; |           | minor    |                          |                                                      |      |
| [#14914](https://github.com/pingcap/tidb/issues/14914)         |           | minor    | <sub>@tiancaiamao</sub>  |                                                      |      |


<h2 name="sig/DDL">sig/DDL</h2>

|                             ISSUE                              |                    CHALLENGE                     | PRIORITY |        ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|--------------------------------------------------|----------|-------------------------|------------------------------------------------------|------|
| [#22117](https://github.com/pingcap/tidb/issues/22117)&#x2757; |                                                  | critical |                         |                                                      |      |
| [#22453](https://github.com/pingcap/tidb/issues/22453)         |                                                  | major    | @tangenta               | [#22458](https://github.com/pingcap/tidb/pull/22458) |      |
| [#22145](https://github.com/pingcap/tidb/issues/22145)         |                                                  | major    | @morgo                  |                                                      |      |
| [#17808](https://github.com/pingcap/tidb/issues/17808)         |                                                  | major    | <sub>@Reminiscent</sub> | [#21497](https://github.com/pingcap/tidb/pull/21497) |      |
| [#11952](https://github.com/pingcap/tidb/issues/11952)         |                                                  | major    | @bb7133                 | [#21564](https://github.com/pingcap/tidb/pull/21564) |      |
| [#982](https://github.com/pingcap/tidb/issues/982)             |                                                  | major    |                         | [#20708](https://github.com/pingcap/tidb/pull/20708) |      |
| [#21943](https://github.com/pingcap/tidb/issues/21943)         |                                                  | moderate | @bb7133                 | [#21953](https://github.com/pingcap/tidb/pull/21953) |      |
| [#21835](https://github.com/pingcap/tidb/issues/21835)         |                                                  | moderate |                         | [#21845](https://github.com/pingcap/tidb/pull/21845) |      |
| [#21063](https://github.com/pingcap/tidb/issues/21063)         |                                                  | moderate |                         | [#21064](https://github.com/pingcap/tidb/pull/21064) |      |
| [#20592](https://github.com/pingcap/tidb/issues/20592)&#x2757; |                                                  | moderate |                         |                                                      |      |
| [#19697](https://github.com/pingcap/tidb/issues/19697)         |                                                  | moderate | <sub>@xiongjiwei</sub>  |                                                      |      |
| [#18907](https://github.com/pingcap/tidb/issues/18907)&#x2757; |                                                  | moderate |                         |                                                      |      |
| [#18336](https://github.com/pingcap/tidb/issues/18336)         |                                                  | moderate | @AilinKid               |                                                      |      |
| [#17745](https://github.com/pingcap/tidb/issues/17745)&#x2757; |                                                  | moderate |                         |                                                      |      |
| [#17686](https://github.com/pingcap/tidb/issues/17686)         |                                                  | moderate | @qw4990                 |                                                      |      |
| [#15567](https://github.com/pingcap/tidb/issues/15567)&#x2757; |                                                  | moderate |                         |                                                      |      |
| [#11648](https://github.com/pingcap/tidb/issues/11648)         |                                                  | moderate |                         | [#21237](https://github.com/pingcap/tidb/pull/21237) |      |
| [#10260](https://github.com/pingcap/tidb/issues/10260)&#x2757; |                                                  | moderate |                         |                                                      |      |
| [#7545](https://github.com/pingcap/tidb/issues/7545)&#x2757;   |                                                  | moderate |                         |                                                      |      |
| [#3876](https://github.com/pingcap/tidb/issues/3876)&#x2757;   |                                                  | moderate |                         |                                                      |      |
| [#3644](https://github.com/pingcap/tidb/issues/3644)&#x2757;   | &#x2665; yes!</br>Mentor: @qw4990</br>Score: 300 | moderate |                         |                                                      |      |
| [#17460](https://github.com/pingcap/tidb/issues/17460)         |                                                  | minor    | @zimulala               |                                                      |      |
| [#14241](https://github.com/pingcap/tidb/issues/14241)&#x2757; |                                                  | minor    |                         |                                                      |      |
| [#11410](https://github.com/pingcap/tidb/issues/11410)&#x2757; |                                                  | minor    |                         |                                                      |      |



---

updated at 2021-01-27T19:21:31-00:00


