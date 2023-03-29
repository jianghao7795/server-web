package frontend

import (
	"log"
	loginRequest "server/model/frontend/request"

	"github.com/gin-gonic/gin"
)

type FrontendUser struct{}

func (u *FrontendUser) Login(c *gin.Context) {
	var user loginRequest.LoginForm
	_ = c.ShouldBindJSON(&user)
	userInfo, err := frontendService.Login(user)
	log.Println(userInfo, err)
}
