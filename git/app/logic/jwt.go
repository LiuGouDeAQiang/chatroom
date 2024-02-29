package logic

import (
	model "chatroom/app/models"
	"chatroom/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var Clients *Client
var clients = make(map[int]*Client)
var Uid int

type Client struct {
	conn *websocket.Conn
	uid  int
}

func Validate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		// 例如返回错误响应
		context.JSON(http.StatusBadRequest, tools.ECode{
			Code:    101,
			Message: "未接收到权限密钥",
			Data:    nil,
			Token:   nil,
		})
		return
	}
	jwt, err := model.CheckJwt(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    102,
			Message: "校验失败",
			Data:    err,
		})
		return
	}
	newUUID := uuid.New()

	Uid = int(jwt.Uid)
	fmt.Println(newUUID.String())
	context.JSON(200, tools.ECode{
		Code:    0,
		Message: "",
		Data:    Uid,
		Token:   nil,
		UUid:    newUUID.String(),
	})
	return
}
