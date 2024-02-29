package logic

import (
	model "chatroom/app/models"
	"chatroom/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGroupList(context *gin.Context) {
	token := context.Query("token")
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
	uid := jwt.Uid
	fmt.Println(" 群列表测试", +uid)
	groupUIDs, err := model.GetGroupList(uid)
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	fmt.Println(groupUIDs)
	context.JSON(200, gin.H{"data": groupUIDs})
}

func GetFriendsList(context *gin.Context) {
	token := context.Query("token")
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
	uid := jwt.Uid
	fmt.Println(" 好友列表测试", +uid)
	FriendsUIDs, err := model.GetFriendsList(uid)
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	fmt.Println(FriendsUIDs)
	context.JSON(200, gin.H{"data": FriendsUIDs})
}
func GetRecentlyList(context *gin.Context) {
	token := context.Query("token")
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
	uid := jwt.Uid
	fmt.Println(" 群列表测试", +uid)
	groupUIDs, err := model.GetGroupList(uid)
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	fmt.Println(groupUIDs)
	context.JSON(200, gin.H{"data": groupUIDs})
}
