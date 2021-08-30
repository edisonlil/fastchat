package control

import (
	"fastchat/domain"
	"fastchat/service"
	"github.com/gin-gonic/gin"
)

//Login 用户登录
func Login(c *gin.Context) {

	user := domain.User{}
	c.BindJSON(&user)

	result := service.UserLogin(user)

	//返回结果
	c.JSON(result.Code, result)
}
