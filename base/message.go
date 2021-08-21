package base

//Message FastChat 消息实体
type Message struct {
	Id string //消息ID 唯一标识

	Data []byte //内容

	MsgLength uint16 //消息长度

	SenderAppId string //发送者Id

	SenderOpenId string //发送者平台ID

	SendTime uint64 //发送时间

	AcceptTime uint64 //消息接受时间

	AcceptAppId string //接收者平台ID

	AcceptOpenId string //接收者平台用户ID

}
