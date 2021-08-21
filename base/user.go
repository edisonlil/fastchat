package base

//User fastChat 用户实体
type User struct {
	Id string //主键

	AppId string //平台唯一标识

	OpenId string //平台用户唯一标识

}

//NewUser 创建用户实体
func NewUser(AppId string, OpenId string) *User {

	return &User{
		AppId:  AppId,
		OpenId: OpenId,
	}
}
