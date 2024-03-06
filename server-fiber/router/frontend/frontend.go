package frontend

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router fiber.Router) {
	frontend := Router.Group("")
	var frontendTagApi = v1.ApiGroupApp.FrontendApiGroup.FrontendTagApi
	{
		frontend.Get("GetTagList", frontendTagApi.GetTagList)
		frontend.Get("GetTagArticleList/:id", frontendTagApi.GetTag)
	}
	var frontendArticleApi = v1.ApiGroupApp.FrontendApiGroup.FrontendArticleApi
	{
		frontend.Get("GetArticleList", frontendArticleApi.GetArticleList)
		frontend.Get("GetArticle/:id", frontendArticleApi.GetArticleDetail)
		frontend.Get("GetSearchArticle/:name/:value", frontendArticleApi.GetSearchArticle)
	}
	var frontendCommentApi = v1.ApiGroupApp.FrontendApiGroup.CommentApi
	{
		frontend.Get("GetArticleComment/:articleId", frontendCommentApi.GetCommentByArticleId)
		frontend.Post("createdComment", frontendCommentApi.CreatedComment)
	}
	var frontendUserApi = v1.ApiGroupApp.FrontendApiGroup.FrontendUser
	{
		frontend.Get("GetImages", frontendUserApi.GetImages)
		frontend.Post("login", frontendUserApi.Login)
		frontend.Get("GetCurrentUser", frontendUserApi.GetCurrent)
		frontend.Put("updateBackgroundImage", frontendUserApi.UpdateUserBackgroudImage)
		frontend.Put("resetPassword", frontendUserApi.UpdatePassword)
		frontend.Post("register", frontendUserApi.RegisterUser)
		frontend.Put("updateUser", frontendUserApi.UpdateUser)
	}
	var frontendUploadApi = v1.ApiGroupApp.AppApiGroup.FileUploadAndDownloadApi
	{
		frontend.Post("upload", frontendUploadApi.UploadFile)
	}
}
