package service

import (
	"fastchat/auth"
	"fastchat/base"
	"fastchat/domain"
	"fastchat/store"
	log "github.com/sirupsen/logrus"
	"time"
)

const MongoColl = "user"

//UserLogin 用户登录
func UserLogin(user *domain.User) *base.Result {

	res := UserRegister(user)

	if !res.Success {
		return res
	}
	data := res.Data.(domain.User)

	token, err := auth.CreateJwtToken(auth.JwtClaims{
		UserId:    data.Id,
		OpenId:    data.OpenId,
		Namespace: data.Namespace,
		Exp:       time.Now().Add(2 * time.Hour).Unix(), // Jwt Token 两小时后过期
	})

	if err != nil {
		return base.ResultFail().SetMsg("Create Token Error：" + err.Error())
	}

	//返回给客户端信息
	return base.ResultSuccess().SetData(map[string]interface{}{
		"Token": token,
		"Users": data,
	})
}

func UserRegister(user *domain.User) *base.Result {

	data := GetUserByNameSpaceAndOpenId(user.Namespace, user.OpenId)

	if data != nil {
		return base.ResultSuccess().SetData(data)
	}

	res, err := store.InsertOne(MongoColl, user)

	if err != nil {
		return base.ResultFail().SetMsg(err.Error())
	}

	return base.ResultSuccess().SetData(res.(domain.User))
}

//GetUserByNameSpaceAndOpenId 获取指定命名空间的OpenId用户
func GetUserByNameSpaceAndOpenId(namespace string, openId string) *domain.User {

	user := &domain.User{}

	err := store.FindOne(MongoColl, map[string]interface{}{
		"namespace": namespace,
		"openId":    openId,
	}, user)

	if err != nil {
		log.Error(err.Error())
	}

	return user
}

func GetUserById(id string) *domain.User {
	user := &domain.User{}

	err := store.FindOneById(MongoColl, id, user)

	if err != nil {
		log.Error(err.Error())
	}

	return user
}
