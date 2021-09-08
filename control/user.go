package control

import (
	"fastchat/domain"
	"fastchat/service"
	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
// @Router /user/login [post]
func UserLogin(c *gin.Context) {

	user := &domain.User{}
	c.BindJSON(user)

	result := service.UserLogin(user)

	//返回结果
	c.JSON(result.Code, result)
}
