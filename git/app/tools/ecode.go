package tools

import "fmt"

var (
	OK       = ECode{Code: 0}
	NotLogin = ECode{Code: 10001, Message: "用户未登录"}
	UserErr  = ECode{Code: 10003, Message: "账号密码错误"}
	ParamErr = ECode{Code: 10002, Message: "参数错误"}
)

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Token   any    `json:"token"`
	UUid    any    `json:"uuid"`
}

func (e *ECode) String() string {
	return fmt.Sprintf("code:%d,message:%s", e.Code, e.Message)
}
