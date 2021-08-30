package base

//响应Code
const (
	SUCCESS = 200

	FAIL = 500
)

//Result Gin Http 返回实体
type Result struct {
	Code int

	Msg string

	Data interface{}
}

func ResultSuccess(msg string) *Result {

	return &Result{
		Code: SUCCESS,
		Msg:  msg,
	}
}

func ResultFail(msg string) *Result {

	return &Result{
		Code: FAIL,
		Msg:  msg,
	}
}

func (p *Result) SetMsg(msg string) *Result {
	p.Msg = msg
	return p
}
