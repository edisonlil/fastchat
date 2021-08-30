package service

import (
	"fastchat/base"
	"fastchat/domain"
	"fastchat/store"
	log "github.com/sirupsen/logrus"
)

const MongoColl = "user"

func UserLogin(user domain.User) *base.Result {

	data := GetUserByNameSpaceAndOpenId(user.Namespace, user.OpenId)

	if data != nil {
		return base.ResultSuccess("")
	}

	_, err := store.InsertOne(MongoColl, user)

	if err != nil {
		return base.ResultFail(err.Error())
	}

	return base.ResultSuccess("")
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
