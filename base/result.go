package base

//响应Code
const (
	SUCCESS = 200

	FAIL = 500
)

//Result Gin Http 返回实体
type Result struct {
	Code int

	Success bool

	Msg string

	Data interface{}
}

func ResultSuccess() *Result {

	return &Result{
		Code:    SUCCESS,
		Success: true,
	}
}

func ResultFail() *Result {

	return &Result{
		Code:    FAIL,
		Success: false,
	}
}

func (p *Result) SetMsg(msg string) *Result {
	p.Msg = msg
	return p
}

func (p *Result) SetData(data interface{}) *Result {
	p.Data = data
	return p
}
