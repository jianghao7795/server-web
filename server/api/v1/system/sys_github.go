package system

import (
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// [
//   {
//     "url": "https://api.github.com/repos/octocat/Hello-World/commits/6dcb09b5b57875f334f61aebed695e2e4193db5e",
//     "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e",
//     "node_id": "MDY6Q29tbWl0NmRjYjA5YjViNTc4NzVmMzM0ZjYxYWViZWQ2OTVlMmU0MTkzZGI1ZQ==",
//     "html_url": "https://github.com/octocat/Hello-World/commit/6dcb09b5b57875f334f61aebed695e2e4193db5e",
//     "comments_url": "https://api.github.com/repos/octocat/Hello-World/commits/6dcb09b5b57875f334f61aebed695e2e4193db5e/comments",
//     "commit": {
//       "url": "https://api.github.com/repos/octocat/Hello-World/git/commits/6dcb09b5b57875f334f61aebed695e2e4193db5e",
//       "author": {
//         "name": "Monalisa Octocat",
//         "email": "support@github.com",
//         "date": "2011-04-14T16:00:49Z"
//       },
//       "committer": {
//         "name": "Monalisa Octocat",
//         "email": "support@github.com",
//         "date": "2011-04-14T16:00:49Z"
//       },
//       "message": "Fix all the bugs",
//       "tree": {
//         "url": "https://api.github.com/repos/octocat/Hello-World/tree/6dcb09b5b57875f334f61aebed695e2e4193db5e",
//         "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e"
//       },
//       "comment_count": 0,
//       "verification": {
//         "verified": false,
//         "reason": "unsigned",
//         "signature": null,
//         "payload": null
//       }
//     },
//     "author": {
//       "login": "octocat",
//       "id": 1,
//       "node_id": "MDQ6VXNlcjE=",
//       "avatar_url": "https://github.com/images/error/octocat_happy.gif",
//       "gravatar_id": "",
//       "url": "https://api.github.com/users/octocat",
//       "html_url": "https://github.com/octocat",
//       "followers_url": "https://api.github.com/users/octocat/followers",
//       "following_url": "https://api.github.com/users/octocat/following{/other_user}",
//       "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
//       "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
//       "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
//       "organizations_url": "https://api.github.com/users/octocat/orgs",
//       "repos_url": "https://api.github.com/users/octocat/repos",
//       "events_url": "https://api.github.com/users/octocat/events{/privacy}",
//       "received_events_url": "https://api.github.com/users/octocat/received_events",
//       "type": "User",
//       "site_admin": false
//     },
//     "committer": {
//       "login": "octocat",
//       "id": 1,
//       "node_id": "MDQ6VXNlcjE=",
//       "avatar_url": "https://github.com/images/error/octocat_happy.gif",
//       "gravatar_id": "",
//       "url": "https://api.github.com/users/octocat",
//       "html_url": "https://github.com/octocat",
//       "followers_url": "https://api.github.com/users/octocat/followers",
//       "following_url": "https://api.github.com/users/octocat/following{/other_user}",
//       "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
//       "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
//       "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
//       "organizations_url": "https://api.github.com/users/octocat/orgs",
//       "repos_url": "https://api.github.com/users/octocat/repos",
//       "events_url": "https://api.github.com/users/octocat/events{/privacy}",
//       "received_events_url": "https://api.github.com/users/octocat/received_events",
//       "type": "User",
//       "site_admin": false
//     },
//     "parents": [
//       {
//         "url": "https://api.github.com/repos/octocat/Hello-World/commits/6dcb09b5b57875f334f61aebed695e2e4193db5e",
//         "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e"
//       }
//     ]
//   }
// ]

type SystemGithubApi struct{}

// type commit struct {
// 	url string
// 	sha string
// 	node_id string
// 	html_url string
// 	comments_url string
// }

var githubService = service.ServiceGroupApp.SystemServiceGroup.GithubService

func (g *SystemGithubApi) GetGithubList(c *gin.Context) {
	var searchInfo request.PageInfo
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	searchInfo.Page, _ = strconv.Atoi(page)
	searchInfo.PageSize, _ = strconv.Atoi(pageSize)

	if list, err := githubService.GetGithubList(searchInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, "获取成功", c)
	}
}

func (g *SystemGithubApi) CreateGithub(c *gin.Context) {
	var data []system.SysGithub
	_ = c.ShouldBindJSON(&data)
	// log.Println(data)
	if err := githubService.CreateApi(data); err != nil {
		global.LOG.Error("创建commit有错误!", zap.Error(err))
		response.FailWithMessage("创建commit有错误!", c)
	} else {
		response.OkWithData(gin.H{}, c)
	}

	// page := "1"
	// per_page := "5"
	// resp, err := http.Get("https://api.github.com/repos/JiangHaoCode/server-web/commits?page=" + page + "&per_page=" + per_page)
	// if err != nil {
	// 	fmt.Printf("get failed, err:%v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("read from resp.Body failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Print(string(body))
}
