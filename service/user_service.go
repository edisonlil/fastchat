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

func UserLogin(user domain.User) *base.Result {

	res := UserRegister(user)

	if !res.Success {
		return res
	}

	data := res.Data.(domain.User)
	auth.CreateJwtToken(auth.JwtClaims{
		UserId:    data.Id,
		Namespace: data.Namespace,
		Exp:       time.Now().Unix(), //TODO 需改为两小时后
	})

	return res
}

func UserRegister(user domain.User) *base.Result {

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

	user := domain.User{}

	err := store.FindOne(MongoColl, map[string]interface{}{
		"namespace": namespace,
		"openId":    openId,
	}, user)

	if err != nil {
		log.Error(err.Error())
	}

	return &user
}
