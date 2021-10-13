package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	Online  = 1
	Offline = 0
)

//User fastChat 用户实体
type User struct {
  
	sx primitive.ObjectID
  
	Id string //主键

	Namespace string //命名空间

	OpenId string //平台用户唯一标识

	Status int //

}

//NewUser 创建用户实体
func NewUser(namespace string, openId string) *User {

	return &User{
		Namespace: namespace,
		OpenId:    openId,
		Status:    Online,
	}
}
