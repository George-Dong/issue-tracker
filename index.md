

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

|                             ISSUE                              | PRIORITY |                              ASSIGNEE                               |                          PR                          | HINT |
|----------------------------------------------------------------|----------|---------------------------------------------------------------------|------------------------------------------------------|------|
| [#8190](https://github.com/pingcap/tidb/issues/8190)           | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#9373](https://github.com/pingcap/tidb/issues/9373)           | moderate | @morgo                                                              |                                                      |      |
| [#10151](https://github.com/pingcap/tidb/issues/10151)&#x2757; | moderate |                                                                     |                                                      |      |
| [#12182](https://github.com/pingcap/tidb/issues/12182)&#x2757; | moderate |                                                                     |                                                      |      |
| [#13167](https://github.com/pingcap/tidb/issues/13167)         | moderate | @fzhedu                                                             |                                                      |      |
| [#13856](https://github.com/pingcap/tidb/issues/13856)         | moderate | @winoros                                                            |                                                      |      |
| [#15514](https://github.com/pingcap/tidb/issues/15514)&#x2757; | minor    |                                                                     |                                                      |      |
| [#16473](https://github.com/pingcap/tidb/issues/16473)         | minor    | @winoros                                                            |                                                      |      |
| [#16764](https://github.com/pingcap/tidb/issues/16764)         | moderate | @lzmhhh123                                                          |                                                      |      |
| [#16788](https://github.com/pingcap/tidb/issues/16788)         | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#16909](https://github.com/pingcap/tidb/issues/16909)         | moderate | @lzmhhh123                                                          |                                                      |      |
| [#17731](https://github.com/pingcap/tidb/issues/17731)         | moderate |                                                                     | [#22416](https://github.com/pingcap/tidb/pull/22416) |      |
| [#17852](https://github.com/pingcap/tidb/issues/17852)         | moderate | @XuHuaiyu                                                           |                                                      |      |
| [#18216](https://github.com/pingcap/tidb/issues/18216)         | moderate | @fzhedu                                                             |                                                      |      |
| [#19713](https://github.com/pingcap/tidb/issues/19713)         | minor    | @winoros                                                            |                                                      |      |
| [#19802](https://github.com/pingcap/tidb/issues/19802)&#x2757; | minor    |                                                                     |                                                      |      |
| [#20019](https://github.com/pingcap/tidb/issues/20019)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21098](https://github.com/pingcap/tidb/issues/21098)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21454](https://github.com/pingcap/tidb/issues/21454)&#x2757; | minor    |                                                                     |                                                      |      |
| [#21475](https://github.com/pingcap/tidb/issues/21475)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21625](https://github.com/pingcap/tidb/issues/21625)&#x2757; | minor    |                                                                     |                                                      |      |
| [#21677](https://github.com/pingcap/tidb/issues/21677)         | major    | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#22082](https://github.com/pingcap/tidb/issues/22082)&#x2757; | moderate |                                                                     |                                                      |      |
| [#22096](https://github.com/pingcap/tidb/issues/22096)&#x2757; | moderate |                                                                     |                                                      |      |
| [#22301](https://github.com/pingcap/tidb/issues/22301)         | minor    | <sub>@Reminiscent</sub></br><sub><sup>@xuyifangreeneyes</sup></sub> |                                                      |      |
| [#22535](https://github.com/pingcap/tidb/issues/22535)&#x2757; | minor    |                                                                     |                                                      |      |
| [#22694](https://github.com/pingcap/tidb/issues/22694)         | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             | [#22722](https://github.com/pingcap/tidb/pull/22722) |      |
| [#22708](https://github.com/pingcap/tidb/issues/22708)         | moderate | @zhuo-zhi                                                           |                                                      |      |
| [#22934](https://github.com/pingcap/tidb/issues/22934)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23391](https://github.com/pingcap/tidb/issues/23391)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23396](https://github.com/pingcap/tidb/issues/23396)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23459](https://github.com/pingcap/tidb/issues/23459)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23499](https://github.com/pingcap/tidb/issues/23499)         | major    | @winoros                                                            |                                                      |      |
| [#23506](https://github.com/pingcap/tidb/issues/23506)         | major    | <sub>@xiongjiwei</sub>                                              | [#26964](https://github.com/pingcap/tidb/pull/26964) |      |
| [#23539](https://github.com/pingcap/tidb/issues/23539)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23582](https://github.com/pingcap/tidb/issues/23582)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23633](https://github.com/pingcap/tidb/issues/23633)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23907](https://github.com/pingcap/tidb/issues/23907)&#x2757; | major    |                                                                     |                                                      |      |
| [#24010](https://github.com/pingcap/tidb/issues/24010)         | moderate | @qw4990</br>@rebelice                                               |                                                      |      |
| [#24012](https://github.com/pingcap/tidb/issues/24012)         | moderate | @qw4990</br>@rebelice                                               |                                                      |      |
| [#24013](https://github.com/pingcap/tidb/issues/24013)         | moderate | @rebelice                                                           |                                                      |      |
| [#24014](https://github.com/pingcap/tidb/issues/24014)         | moderate | @rebelice                                                           |                                                      |      |
| [#24015](https://github.com/pingcap/tidb/issues/24015)         | moderate | @rebelice                                                           |                                                      |      |
| [#24057](https://github.com/pingcap/tidb/issues/24057)         | moderate | @winoros                                                            |                                                      |      |
| [#24318](https://github.com/pingcap/tidb/issues/24318)         | moderate | @qw4990                                                             |                                                      |      |
| [#24323](https://github.com/pingcap/tidb/issues/24323)&#x2757; | minor    |                                                                     |                                                      |      |
| [#24324](https://github.com/pingcap/tidb/issues/24324)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#24449](https://github.com/pingcap/tidb/issues/24449)         | minor    | @zoomxi                                                             | [#26386](https://github.com/pingcap/tidb/pull/26386) |      |
| [#24452](https://github.com/pingcap/tidb/issues/24452)         | moderate | @winoros                                                            |                                                      |      |
| [#24512](https://github.com/pingcap/tidb/issues/24512)&#x2757; | moderate |                                                                     |                                                      |      |
| [#24550](https://github.com/pingcap/tidb/issues/24550)         | moderate | @eurekaka                                                           |                                                      |      |
| [#24563](https://github.com/pingcap/tidb/issues/24563)         | moderate | @eurekaka                                                           |                                                      |      |
| [#24567](https://github.com/pingcap/tidb/issues/24567)&#x2757; | major    |                                                                     |                                                      |      |
| [#24594](https://github.com/pingcap/tidb/issues/24594)         | major    | @winoros                                                            |                                                      |      |
| [#24622](https://github.com/pingcap/tidb/issues/24622)&#x2757; | moderate |                                                                     |                                                      |      |
| [#24667](https://github.com/pingcap/tidb/issues/24667)         | moderate | <sub><sup>@time-and-fate</sup></sub>                                | [#24669](https://github.com/pingcap/tidb/pull/24669) |      |
| [#24679](https://github.com/pingcap/tidb/issues/24679)         | moderate | @rebelice                                                           |                                                      |      |
| [#24855](https://github.com/pingcap/tidb/issues/24855)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#25043](https://github.com/pingcap/tidb/issues/25043)&#x2757; | major    |                                                                     |                                                      |      |
| [#25066](https://github.com/pingcap/tidb/issues/25066)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#25086](https://github.com/pingcap/tidb/issues/25086)&#x2757; | minor    |                                                                     |                                                      |      |
| [#25144](https://github.com/pingcap/tidb/issues/25144)         | major    | @rebelice                                                           |                                                      |      |
| [#25239](https://github.com/pingcap/tidb/issues/25239)         | moderate | @rebelice                                                           |                                                      |      |
| [#25364](https://github.com/pingcap/tidb/issues/25364)         | major    |                                                                     | [#25390](https://github.com/pingcap/tidb/pull/25390) |      |
| [#25392](https://github.com/pingcap/tidb/issues/25392)         | major    | @eurekaka                                                           |                                                      |      |
| [#25422](https://github.com/pingcap/tidb/issues/25422)&#x2757; | major    |                                                                     |                                                      |      |
| [#25490](https://github.com/pingcap/tidb/issues/25490)&#x2757; | major    |                                                                     |                                                      |      |
| [#25539](https://github.com/pingcap/tidb/issues/25539)&#x2757; | moderate |                                                                     |                                                      |      |
| [#25585](https://github.com/pingcap/tidb/issues/25585)&#x2757; | major    |                                                                     |                                                      |      |
| [#25600](https://github.com/pingcap/tidb/issues/25600)&#x2757; | moderate |                                                                     |                                                      |      |
| [#25603](https://github.com/pingcap/tidb/issues/25603)&#x2757; | moderate |                                                                     |                                                      |      |
| [#25646](https://github.com/pingcap/tidb/issues/25646)&#x2757; | major    |                                                                     |                                                      |      |
| [#25727](https://github.com/pingcap/tidb/issues/25727)         | major    | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#25782](https://github.com/pingcap/tidb/issues/25782)         | major    | @winoros                                                            |                                                      |      |
| [#25812](https://github.com/pingcap/tidb/issues/25812)         | major    | @winoros                                                            |                                                      |      |
| [#25852](https://github.com/pingcap/tidb/issues/25852)         | moderate | @qw4990                                                             |                                                      |      |
| [#26077](https://github.com/pingcap/tidb/issues/26077)         | major    | <sub>@Reminiscent</sub>                                             |                                                      |      |
| [#26166](https://github.com/pingcap/tidb/issues/26166)         | major    | @bb7133</br>@mjonss                                                 |                                                      |      |
| [#26249](https://github.com/pingcap/tidb/issues/26249)&#x2757; | moderate |                                                                     |                                                      |      |
| [#26377](https://github.com/pingcap/tidb/issues/26377)         | minor    | <sub>@tiancaiamao</sub>                                             | [#27347](https://github.com/pingcap/tidb/pull/27347) |      |
| [#26547](https://github.com/pingcap/tidb/issues/26547)         | moderate | @winoros                                                            | [#27359](https://github.com/pingcap/tidb/pull/27359) |      |
| [#26569](https://github.com/pingcap/tidb/issues/26569)         | critical | @winoros                                                            |                                                      |      |
| [#26576](https://github.com/pingcap/tidb/issues/26576)&#x2757; | moderate |                                                                     |                                                      |      |
| [#26638](https://github.com/pingcap/tidb/issues/26638)         | major    | <sub>@Reminiscent</sub>                                             | [#26713](https://github.com/pingcap/tidb/pull/26713) |      |
| [#26754](https://github.com/pingcap/tidb/issues/26754)         | moderate | @qw4990                                                             |                                                      |      |
| [#26764](https://github.com/pingcap/tidb/issues/26764)         | moderate | @winoros                                                            |                                                      |      |
| [#26779](https://github.com/pingcap/tidb/issues/26779)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#26782](https://github.com/pingcap/tidb/issues/26782)         | critical | <sub>@guo-shaoge</sub>                                              |                                                      |      |
| [#26873](https://github.com/pingcap/tidb/issues/26873)         | major    | <sub>@Reminiscent</sub>                                             |                                                      |      |
| [#26944](https://github.com/pingcap/tidb/issues/26944)         |          | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#26945](https://github.com/pingcap/tidb/issues/26945)         | major    | <sub>@hawkingrei</sub>                                              |                                                      |      |
| [#26950](https://github.com/pingcap/tidb/issues/26950)         | moderate | @qw4990                                                             |                                                      |      |
| [#27070](https://github.com/pingcap/tidb/issues/27070)         | minor    | @qw4990                                                             |                                                      |      |
| [#27083](https://github.com/pingcap/tidb/issues/27083)         | major    | <sub><sup>@xuyifangreeneyes</sup></sub>                             | [#27161](https://github.com/pingcap/tidb/pull/27161) |      |
| [#27093](https://github.com/pingcap/tidb/issues/27093)         |          | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#27106](https://github.com/pingcap/tidb/issues/27106)         | critical | @rebelice                                                           |                                                      |      |
| [#27159](https://github.com/pingcap/tidb/issues/27159)         | critical | <sub>@wjhuang2016</sub>                                             | [#27170](https://github.com/pingcap/tidb/pull/27170) |      |
| [#27166](https://github.com/pingcap/tidb/issues/27166)         | moderate | @chrysan                                                            |                                                      |      |
| [#27187](https://github.com/pingcap/tidb/issues/27187)         | critical | @chrysan                                                            | [#27359](https://github.com/pingcap/tidb/pull/27359) |      |
| [#27241](https://github.com/pingcap/tidb/issues/27241)         | moderate | @qw4990                                                             |                                                      |      |
| [#27249](https://github.com/pingcap/tidb/issues/27249)         | moderate | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#27272](https://github.com/pingcap/tidb/issues/27272)         | major    |                                                                     | [#27275](https://github.com/pingcap/tidb/pull/27275) |      |
| [#27313](https://github.com/pingcap/tidb/issues/27313)         | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#27328](https://github.com/pingcap/tidb/issues/27328)         | major    | @qw4990                                                             |                                                      |      |
| [#27346](https://github.com/pingcap/tidb/issues/27346)         | moderate | @qw4990                                                             |                                                      |      |
| [#27384](https://github.com/pingcap/tidb/issues/27384)         | moderate | @mjonss</br>@qw4990                                                 |                                                      |      |


<h2 name="sig/execution">sig/execution</h2>

|                             ISSUE                              | PRIORITY |        ASSIGNEE         |                                                      PR                                                       | HINT |
|----------------------------------------------------------------|----------|-------------------------|---------------------------------------------------------------------------------------------------------------|------|
| [#8205](https://github.com/pingcap/tidb/issues/8205)&#x2757;   | minor    |                         |                                                                                                               |      |
| [#11866](https://github.com/pingcap/tidb/issues/11866)         | moderate | @dragonly               |                                                                                                               |      |
| [#11932](https://github.com/pingcap/tidb/issues/11932)&#x2757; | moderate |                         |                                                                                                               |      |
| [#13018](https://github.com/pingcap/tidb/issues/13018)         | minor    | @qw4990                 |                                                                                                               |      |
| [#13136](https://github.com/pingcap/tidb/issues/13136)&#x2757; | moderate |                         |                                                                                                               |      |
| [#13157](https://github.com/pingcap/tidb/issues/13157)&#x2757; | moderate |                         |                                                                                                               |      |
| [#13440](https://github.com/pingcap/tidb/issues/13440)         | minor    | <sub>@SunRunAway</sub>  |                                                                                                               |      |
| [#14399](https://github.com/pingcap/tidb/issues/14399)&#x2757; | moderate |                         |                                                                                                               |      |
| [#15234](https://github.com/pingcap/tidb/issues/15234)         | moderate |                         | [#20015](https://github.com/pingcap/tidb/pull/20015)                                                          |      |
| [#15608](https://github.com/pingcap/tidb/issues/15608)         | moderate | @ichn-hu                |                                                                                                               |      |
| [#15884](https://github.com/pingcap/tidb/issues/15884)&#x2757; | moderate |                         |                                                                                                               |      |
| [#17083](https://github.com/pingcap/tidb/issues/17083)         | moderate | @mmyj                   |                                                                                                               |      |
| [#17489](https://github.com/pingcap/tidb/issues/17489)&#x2757; | minor    |                         |                                                                                                               |      |
| [#17677](https://github.com/pingcap/tidb/issues/17677)&#x2757; | moderate |                         |                                                                                                               |      |
| [#17751](https://github.com/pingcap/tidb/issues/17751)&#x2757; | minor    |                         |                                                                                                               |      |
| [#17832](https://github.com/pingcap/tidb/issues/17832)         | minor    | @qw4990                 |                                                                                                               |      |
| [#17993](https://github.com/pingcap/tidb/issues/17993)&#x2757; | minor    |                         |                                                                                                               |      |
| [#18488](https://github.com/pingcap/tidb/issues/18488)         | moderate | @morgo                  |                                                                                                               |      |
| [#18493](https://github.com/pingcap/tidb/issues/18493)         | minor    | @fzhedu                 |                                                                                                               |      |
| [#19025](https://github.com/pingcap/tidb/issues/19025)         | moderate | <sub>@SunRunAway</sub>  |                                                                                                               |      |
| [#20411](https://github.com/pingcap/tidb/issues/20411)         | moderate | @qw4990                 |                                                                                                               |      |
| [#20563](https://github.com/pingcap/tidb/issues/20563)         | moderate | <sub>@wjhuang2016</sub> | [#1129](https://github.com/pingcap/parser/pull/1129)                                                          |      |
| [#21307](https://github.com/pingcap/tidb/issues/21307)&#x2757; | minor    |                         |                                                                                                               |      |
| [#21584](https://github.com/pingcap/tidb/issues/21584)&#x2757; | major    |                         |                                                                                                               |      |
| [#21653](https://github.com/pingcap/tidb/issues/21653)         | moderate |                         | [#21230](https://github.com/pingcap/tidb/pull/21230)                                                          |      |
| [#21787](https://github.com/pingcap/tidb/issues/21787)&#x2757; | minor    |                         |                                                                                                               |      |
| [#22088](https://github.com/pingcap/tidb/issues/22088)&#x2757; | moderate |                         |                                                                                                               |      |
| [#22132](https://github.com/pingcap/tidb/issues/22132)         | major    | @wshwsh12               | [#22347](https://github.com/pingcap/tidb/pull/22347)                                                          |      |
| [#22206](https://github.com/pingcap/tidb/issues/22206)         | moderate |                         | [#22616](https://github.com/pingcap/tidb/pull/22616)                                                          |      |
| [#22227](https://github.com/pingcap/tidb/issues/22227)&#x2757; | minor    |                         |                                                                                                               |      |
| [#22386](https://github.com/pingcap/tidb/issues/22386)&#x2757; | moderate |                         |                                                                                                               |      |
| [#22394](https://github.com/pingcap/tidb/issues/22394)         | moderate |                         | [#22407](https://github.com/pingcap/tidb/pull/22407)                                                          |      |
| [#22399](https://github.com/pingcap/tidb/issues/22399)&#x2757; | minor    |                         |                                                                                                               |      |
| [#22423](https://github.com/pingcap/tidb/issues/22423)&#x2757; | minor    |                         |                                                                                                               |      |
| [#22525](https://github.com/pingcap/tidb/issues/22525)         | major    | @lzmhhh123              |                                                                                                               |      |
| [#22592](https://github.com/pingcap/tidb/issues/22592)&#x2757; | major    |                         |                                                                                                               |      |
| [#22598](https://github.com/pingcap/tidb/issues/22598)&#x2757; | moderate |                         |                                                                                                               |      |
| [#22604](https://github.com/pingcap/tidb/issues/22604)&#x2757; | major    |                         |                                                                                                               |      |
| [#22665](https://github.com/pingcap/tidb/issues/22665)         | major    |                         | [#22666](https://github.com/pingcap/tidb/pull/22666)                                                          |      |
| [#22735](https://github.com/pingcap/tidb/issues/22735)&#x2757; | critical |                         |                                                                                                               |      |
| [#22749](https://github.com/pingcap/tidb/issues/22749)&#x2757; | minor    |                         |                                                                                                               |      |
| [#22791](https://github.com/pingcap/tidb/issues/22791)         | major    | @johan-j                | [#22823](https://github.com/pingcap/tidb/pull/22823)                                                          |      |
| [#23101](https://github.com/pingcap/tidb/issues/23101)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23110](https://github.com/pingcap/tidb/issues/23110)         | moderate | @jyz0309                |                                                                                                               |      |
| [#23159](https://github.com/pingcap/tidb/issues/23159)         | moderate |                         | [#23206](https://github.com/pingcap/tidb/pull/23206)                                                          |      |
| [#23344](https://github.com/pingcap/tidb/issues/23344)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23366](https://github.com/pingcap/tidb/issues/23366)&#x2757; | minor    |                         |                                                                                                               |      |
| [#23387](https://github.com/pingcap/tidb/issues/23387)         | major    | <sub>@tiancaiamao</sub> | [#23403](https://github.com/pingcap/tidb/pull/23403)                                                          |      |
| [#23411](https://github.com/pingcap/tidb/issues/23411)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23501](https://github.com/pingcap/tidb/issues/23501)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23508](https://github.com/pingcap/tidb/issues/23508)         | major    | <sub>@wjhuang2016</sub> | [#23559](https://github.com/pingcap/tidb/pull/23559)                                                          |      |
| [#23512](https://github.com/pingcap/tidb/issues/23512)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23531](https://github.com/pingcap/tidb/issues/23531)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23552](https://github.com/pingcap/tidb/issues/23552)&#x2757; | moderate |                         |                                                                                                               |      |
| [#23865](https://github.com/pingcap/tidb/issues/23865)         | moderate | @xhebox                 |                                                                                                               |      |
| [#23869](https://github.com/pingcap/tidb/issues/23869)         | moderate | @ichn-hu                |                                                                                                               |      |
| [#23897](https://github.com/pingcap/tidb/issues/23897)         | major    | @wshwsh12               |                                                                                                               |      |
| [#23898](https://github.com/pingcap/tidb/issues/23898)         | major    | @lzmhhh123              |                                                                                                               |      |
| [#23952](https://github.com/pingcap/tidb/issues/23952)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24044](https://github.com/pingcap/tidb/issues/24044)         | minor    | @zoomxi                 | [#26324](https://github.com/pingcap/tidb/pull/26324)                                                          |      |
| [#24134](https://github.com/pingcap/tidb/issues/24134)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24227](https://github.com/pingcap/tidb/issues/24227)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24271](https://github.com/pingcap/tidb/issues/24271)&#x2757; | minor    |                         |                                                                                                               |      |
| [#24284](https://github.com/pingcap/tidb/issues/24284)&#x2757; | minor    |                         |                                                                                                               |      |
| [#24319](https://github.com/pingcap/tidb/issues/24319)&#x2757; | minor    |                         |                                                                                                               |      |
| [#24627](https://github.com/pingcap/tidb/issues/24627)&#x2757; | major    |                         |                                                                                                               |      |
| [#24725](https://github.com/pingcap/tidb/issues/24725)         | moderate | @wzru                   |                                                                                                               |      |
| [#24917](https://github.com/pingcap/tidb/issues/24917)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24928](https://github.com/pingcap/tidb/issues/24928)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24969](https://github.com/pingcap/tidb/issues/24969)&#x2757; | moderate |                         |                                                                                                               |      |
| [#24997](https://github.com/pingcap/tidb/issues/24997)         | major    | @XuHuaiyu               | [#26726](https://github.com/pingcap/tidb/pull/26726)                                                          |      |
| [#25020](https://github.com/pingcap/tidb/issues/25020)         | major    | @zimulala               |                                                                                                               |      |
| [#25032](https://github.com/pingcap/tidb/issues/25032)&#x2757; | major    |                         |                                                                                                               |      |
| [#25053](https://github.com/pingcap/tidb/issues/25053)         | moderate | @wzru                   | [#27119](https://github.com/pingcap/tidb/pull/27119)                                                          |      |
| [#25196](https://github.com/pingcap/tidb/issues/25196)&#x2757; | minor    |                         |                                                                                                               |      |
| [#25199](https://github.com/pingcap/tidb/issues/25199)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25217](https://github.com/pingcap/tidb/issues/25217)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25235](https://github.com/pingcap/tidb/issues/25235)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25245](https://github.com/pingcap/tidb/issues/25245)&#x2757; | major    |                         |                                                                                                               |      |
| [#25333](https://github.com/pingcap/tidb/issues/25333)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25482](https://github.com/pingcap/tidb/issues/25482)         | moderate | @xhebox                 |                                                                                                               |      |
| [#25497](https://github.com/pingcap/tidb/issues/25497)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25529](https://github.com/pingcap/tidb/issues/25529)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25579](https://github.com/pingcap/tidb/issues/25579)         | moderate | @bb7133                 |                                                                                                               |      |
| [#25645](https://github.com/pingcap/tidb/issues/25645)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25691](https://github.com/pingcap/tidb/issues/25691)         | major    | @XuHuaiyu               | [#26892](https://github.com/pingcap/tidb/pull/26892)                                                          |      |
| [#25726](https://github.com/pingcap/tidb/issues/25726)&#x2757; | minor    |                         |                                                                                                               |      |
| [#25734](https://github.com/pingcap/tidb/issues/25734)         | major    | @lzmhhh123              |                                                                                                               |      |
| [#25753](https://github.com/pingcap/tidb/issues/25753)&#x2757; | minor    |                         |                                                                                                               |      |
| [#25802](https://github.com/pingcap/tidb/issues/25802)         | major    | <sub>@guo-shaoge</sub>  |                                                                                                               |      |
| [#25813](https://github.com/pingcap/tidb/issues/25813)&#x2757; | major    |                         |                                                                                                               |      |
| [#25829](https://github.com/pingcap/tidb/issues/25829)         | moderate |                         | [#25879](https://github.com/pingcap/tidb/pull/25879)                                                          |      |
| [#25848](https://github.com/pingcap/tidb/issues/25848)&#x2757; | moderate |                         |                                                                                                               |      |
| [#25898](https://github.com/pingcap/tidb/issues/25898)         | major    | @wzru</br>@XuHuaiyu     | [#26892](https://github.com/pingcap/tidb/pull/26892)                                                          |      |
| [#25993](https://github.com/pingcap/tidb/issues/25993)         | moderate |                         | [#26097](https://github.com/pingcap/tidb/pull/26097)                                                          |      |
| [#26004](https://github.com/pingcap/tidb/issues/26004)         | minor    |                         | [#26005](https://github.com/pingcap/tidb/pull/26005)                                                          |      |
| [#26151](https://github.com/pingcap/tidb/issues/26151)         | moderate |                         | [#26152](https://github.com/pingcap/tidb/pull/26152)                                                          |      |
| [#26358](https://github.com/pingcap/tidb/issues/26358)&#x2757; | major    |                         |                                                                                                               |      |
| [#26384](https://github.com/pingcap/tidb/issues/26384)&#x2757; | moderate |                         |                                                                                                               |      |
| [#26402](https://github.com/pingcap/tidb/issues/26402)&#x2757; | moderate |                         |                                                                                                               |      |
| [#26434](https://github.com/pingcap/tidb/issues/26434)&#x2757; | major    |                         |                                                                                                               |      |
| [#26447](https://github.com/pingcap/tidb/issues/26447)&#x2757; | moderate |                         |                                                                                                               |      |
| [#26485](https://github.com/pingcap/tidb/issues/26485)&#x2757; | moderate |                         |                                                                                                               |      |
| [#26539](https://github.com/pingcap/tidb/issues/26539)         | major    | @lzmhhh123              | [#26892](https://github.com/pingcap/tidb/pull/26892)                                                          |      |
| [#26554](https://github.com/pingcap/tidb/issues/26554)         | major    | @wshwsh12               | [#27022](https://github.com/pingcap/tidb/pull/27022)                                                          |      |
| [#26703](https://github.com/pingcap/tidb/issues/26703)&#x2757; | major    |                         |                                                                                                               |      |
| [#26790](https://github.com/pingcap/tidb/issues/26790)         | major    | @wzru                   | [#27005](https://github.com/pingcap/tidb/pull/27005)</br>[#27006](https://github.com/pingcap/tidb/pull/27006) |      |
| [#26806](https://github.com/pingcap/tidb/issues/26806)&#x2757; | major    |                         |                                                                                                               |      |
| [#26885](https://github.com/pingcap/tidb/issues/26885)&#x2757; | major    |                         |                                                                                                               |      |
| [#26886](https://github.com/pingcap/tidb/issues/26886)         | major    | @Yisaer                 |                                                                                                               |      |
| [#26887](https://github.com/pingcap/tidb/issues/26887)         | moderate |                         | [#26888](https://github.com/pingcap/tidb/pull/26888)                                                          |      |
| [#26977](https://github.com/pingcap/tidb/issues/26977)         | major    |                         | [#27122](https://github.com/pingcap/tidb/pull/27122)                                                          |      |
| [#26993](https://github.com/pingcap/tidb/issues/26993)         | major    | <sub>@feitian124</sub>  | [#27128](https://github.com/pingcap/tidb/pull/27128)</br>[#27403](https://github.com/pingcap/tidb/pull/27403) |      |
| [#27078](https://github.com/pingcap/tidb/issues/27078)&#x2757; | major    |                         |                                                                                                               |      |
| [#27135](https://github.com/pingcap/tidb/issues/27135)         | major    | @XuHuaiyu               |                                                                                                               |      |
| [#27232](https://github.com/pingcap/tidb/issues/27232)         | major    | @lzmhhh123              | [#27244](https://github.com/pingcap/tidb/pull/27244)                                                          |      |
| [#27274](https://github.com/pingcap/tidb/issues/27274)&#x2757; |          |                         |                                                                                                               |      |
| [#27296](https://github.com/pingcap/tidb/issues/27296)         | major    |                         | [#27376](https://github.com/pingcap/tidb/pull/27376)                                                          |      |
| [#27386](https://github.com/pingcap/tidb/issues/27386)         | major    | <sub>@wjhuang2016</sub> |                                                                                                               |      |


<h2 name="sig/transaction">sig/transaction</h2>

|                             ISSUE                              | PRIORITY |               ASSIGNEE                |                          PR                          | HINT |
|----------------------------------------------------------------|----------|---------------------------------------|------------------------------------------------------|------|
| [#9762](https://github.com/pingcap/tidb/issues/9762)&#x2757;   | major    |                                       |                                                      |      |
| [#10657](https://github.com/pingcap/tidb/issues/10657)         | major    | @nolouch                              |                                                      |      |
| [#13958](https://github.com/pingcap/tidb/issues/13958)         | moderate | @fzhedu                               |                                                      |      |
| [#14914](https://github.com/pingcap/tidb/issues/14914)         | minor    | <sub>@tiancaiamao</sub>               |                                                      |      |
| [#18048](https://github.com/pingcap/tidb/issues/18048)         | moderate | @qw4990                               |                                                      |      |
| [#20028](https://github.com/pingcap/tidb/issues/20028)         | major    | <sub>@tiancaiamao</sub>               | [#21148](https://github.com/pingcap/tidb/pull/21148) |      |
| [#20949](https://github.com/pingcap/tidb/issues/20949)&#x2757; | minor    |                                       |                                                      |      |
| [#20990](https://github.com/pingcap/tidb/issues/20990)&#x2757; | moderate |                                       |                                                      |      |
| [#21335](https://github.com/pingcap/tidb/issues/21335)         | minor    | @you06                                | [#22146](https://github.com/pingcap/tidb/pull/22146) |      |
| [#21355](https://github.com/pingcap/tidb/issues/21355)&#x2757; | moderate |                                       |                                                      |      |
| [#21506](https://github.com/pingcap/tidb/issues/21506)         | moderate | @cfzjywxk                             |                                                      |      |
| [#21688](https://github.com/pingcap/tidb/issues/21688)         | moderate | @you06                                |                                                      |      |
| [#22244](https://github.com/pingcap/tidb/issues/22244)&#x2757; | moderate |                                       |                                                      |      |
| [#22345](https://github.com/pingcap/tidb/issues/22345)         | minor    | @lysu                                 | [#22372](https://github.com/pingcap/tidb/pull/22372) |      |
| [#22356](https://github.com/pingcap/tidb/issues/22356)&#x2757; | moderate |                                       |                                                      |      |
| [#22393](https://github.com/pingcap/tidb/issues/22393)&#x2757; | moderate |                                       |                                                      |      |
| [#22783](https://github.com/pingcap/tidb/issues/22783)         | minor    | <sub>@tiancaiamao</sub>               |                                                      |      |
| [#22927](https://github.com/pingcap/tidb/issues/22927)&#x2757; | critical |                                       |                                                      |      |
| [#23015](https://github.com/pingcap/tidb/issues/23015)&#x2757; | minor    |                                       |                                                      |      |
| [#23180](https://github.com/pingcap/tidb/issues/23180)&#x2757; | minor    |                                       |                                                      |      |
| [#23235](https://github.com/pingcap/tidb/issues/23235)         | minor    | @tangenta                             |                                                      |      |
| [#23331](https://github.com/pingcap/tidb/issues/23331)         | major    | @cfzjywxk                             | [#23342](https://github.com/pingcap/tidb/pull/23342) |      |
| [#23423](https://github.com/pingcap/tidb/issues/23423)&#x2757; | moderate |                                       |                                                      |      |
| [#23542](https://github.com/pingcap/tidb/issues/23542)         | minor    |                                       | [#24140](https://github.com/pingcap/tidb/pull/24140) |      |
| [#23709](https://github.com/pingcap/tidb/issues/23709)&#x2757; | moderate |                                       |                                                      |      |
| [#23797](https://github.com/pingcap/tidb/issues/23797)         | major    | @tangenta                             |                                                      |      |
| [#24195](https://github.com/pingcap/tidb/issues/24195)&#x2757; | major    |                                       |                                                      |      |
| [#24411](https://github.com/pingcap/tidb/issues/24411)&#x2757; | minor    |                                       |                                                      |      |
| [#24428](https://github.com/pingcap/tidb/issues/24428)         | major    | <sub>@tiancaiamao</sub>               |                                                      |      |
| [#24589](https://github.com/pingcap/tidb/issues/24589)&#x2757; | moderate |                                       |                                                      |      |
| [#24858](https://github.com/pingcap/tidb/issues/24858)&#x2757; | minor    |                                       |                                                      |      |
| [#25003](https://github.com/pingcap/tidb/issues/25003)         | major    | @cfzjywxk                             |                                                      |      |
| [#25029](https://github.com/pingcap/tidb/issues/25029)&#x2757; | minor    |                                       |                                                      |      |
| [#25176](https://github.com/pingcap/tidb/issues/25176)&#x2757; | moderate |                                       |                                                      |      |
| [#25419](https://github.com/pingcap/tidb/issues/25419)         | major    | @breeswish</br><sub>@crazycs520</sub> |                                                      |      |
| [#25457](https://github.com/pingcap/tidb/issues/25457)&#x2757; | major    |                                       |                                                      |      |
| [#25659](https://github.com/pingcap/tidb/issues/25659)&#x2757; | major    |                                       |                                                      |      |
| [#25809](https://github.com/pingcap/tidb/issues/25809)         | critical | @lysu                                 | [#25905](https://github.com/pingcap/tidb/pull/25905) |      |
| [#25846](https://github.com/pingcap/tidb/issues/25846)&#x2757; | major    |                                       |                                                      |      |
| [#25986](https://github.com/pingcap/tidb/issues/25986)         | major    | @qw4990                               |                                                      |      |
| [#26213](https://github.com/pingcap/tidb/issues/26213)         | major    |                                       | [#26248](https://github.com/pingcap/tidb/pull/26248) |      |
| [#26235](https://github.com/pingcap/tidb/issues/26235)         | major    | @bb7133                               |                                                      |      |
| [#26544](https://github.com/pingcap/tidb/issues/26544)         | major    | @sticnarf                             |                                                      |      |
| [#26548](https://github.com/pingcap/tidb/issues/26548)&#x2757; | minor    |                                       |                                                      |      |
| [#26552](https://github.com/pingcap/tidb/issues/26552)         | major    | @lysu                                 |                                                      |      |
| [#26688](https://github.com/pingcap/tidb/issues/26688)         | major    | @morgo                                |                                                      |      |
| [#26801](https://github.com/pingcap/tidb/issues/26801)&#x2757; | moderate |                                       |                                                      |      |
| [#26805](https://github.com/pingcap/tidb/issues/26805)&#x2757; |          |                                       |                                                      |      |
| [#26810](https://github.com/pingcap/tidb/issues/26810)&#x2757; | moderate |                                       |                                                      |      |
| [#26832](https://github.com/pingcap/tidb/issues/26832)         | major    | @lysu                                 |                                                      |      |
| [#26842](https://github.com/pingcap/tidb/issues/26842)         | minor    |                                       | [#26845](https://github.com/pingcap/tidb/pull/26845) |      |
| [#26999](https://github.com/pingcap/tidb/issues/26999)&#x2757; | major    |                                       |                                                      |      |
| [#27058](https://github.com/pingcap/tidb/issues/27058)&#x2757; | major    |                                       |                                                      |      |
| [#27116](https://github.com/pingcap/tidb/issues/27116)         | major    | <sub>@longfangsong</sub>              | [#27235](https://github.com/pingcap/tidb/pull/27235) |      |
| [#27156](https://github.com/pingcap/tidb/issues/27156)         | major    | @you06                                |                                                      |      |
| [#27197](https://github.com/pingcap/tidb/issues/27197)         |          | <sub>@tiancaiamao</sub>               | [#27209](https://github.com/pingcap/tidb/pull/27209) |      |


<h2 name="sig/DDL">sig/DDL</h2>

|                             ISSUE                              | PRIORITY |        ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|----------|-------------------------|------------------------------------------------------|------|
| [#982](https://github.com/pingcap/tidb/issues/982)             | major    | @tangenta               | [#20708](https://github.com/pingcap/tidb/pull/20708) |      |
| [#3644](https://github.com/pingcap/tidb/issues/3644)&#x2757;   | moderate |                         |                                                      |      |
| [#3876](https://github.com/pingcap/tidb/issues/3876)&#x2757;   | moderate |                         |                                                      |      |
| [#7545](https://github.com/pingcap/tidb/issues/7545)&#x2757;   | moderate |                         |                                                      |      |
| [#10260](https://github.com/pingcap/tidb/issues/10260)&#x2757; | moderate |                         |                                                      |      |
| [#11410](https://github.com/pingcap/tidb/issues/11410)&#x2757; | minor    |                         |                                                      |      |
| [#11648](https://github.com/pingcap/tidb/issues/11648)         | moderate |                         | [#21237](https://github.com/pingcap/tidb/pull/21237) |      |
| [#11952](https://github.com/pingcap/tidb/issues/11952)         | major    | @bb7133                 | [#21564](https://github.com/pingcap/tidb/pull/21564) |      |
| [#14241](https://github.com/pingcap/tidb/issues/14241)&#x2757; | minor    |                         |                                                      |      |
| [#15567](https://github.com/pingcap/tidb/issues/15567)&#x2757; | moderate |                         |                                                      |      |
| [#17460](https://github.com/pingcap/tidb/issues/17460)         | minor    | @zimulala               |                                                      |      |
| [#17686](https://github.com/pingcap/tidb/issues/17686)         | moderate | @qw4990                 |                                                      |      |
| [#17745](https://github.com/pingcap/tidb/issues/17745)&#x2757; | moderate |                         |                                                      |      |
| [#17808](https://github.com/pingcap/tidb/issues/17808)         | major    | <sub>@Reminiscent</sub> | [#21497](https://github.com/pingcap/tidb/pull/21497) |      |
| [#18336](https://github.com/pingcap/tidb/issues/18336)         | moderate | @AilinKid               |                                                      |      |
| [#18907](https://github.com/pingcap/tidb/issues/18907)&#x2757; | moderate |                         |                                                      |      |
| [#19697](https://github.com/pingcap/tidb/issues/19697)         | moderate | <sub>@xiongjiwei</sub>  |                                                      |      |
| [#20592](https://github.com/pingcap/tidb/issues/20592)&#x2757; | moderate |                         |                                                      |      |
| [#21063](https://github.com/pingcap/tidb/issues/21063)         | moderate |                         | [#21064](https://github.com/pingcap/tidb/pull/21064) |      |
| [#21835](https://github.com/pingcap/tidb/issues/21835)         | moderate |                         | [#21845](https://github.com/pingcap/tidb/pull/21845) |      |
| [#21943](https://github.com/pingcap/tidb/issues/21943)         | moderate | @bb7133                 | [#21953](https://github.com/pingcap/tidb/pull/21953) |      |
| [#22117](https://github.com/pingcap/tidb/issues/22117)&#x2757; | critical |                         |                                                      |      |
| [#22453](https://github.com/pingcap/tidb/issues/22453)         | major    | @tangenta               | [#22458](https://github.com/pingcap/tidb/pull/22458) |      |
| [#22704](https://github.com/pingcap/tidb/issues/22704)&#x2757; |          |                         |                                                      |      |



---

updated at 2021-08-20T03:15:10-00:00


