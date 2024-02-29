package router

import (
	"chatroom/app/logic"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()
	r.GET("/ws", logic.HandWebSocket)
	r.POST("/login", logic.PostLogin)
	r.POST("/validate-jwt", logic.Validate)
	r.POST("/create", logic.CreateUser)
	go logic.HandleMessages()
	r.GET("/groupsList", logic.GetGroupList)
	r.GET("/friendsList", logic.GetFriendsList)
	r.GET("/recentlyList", logic.GetRecentlyList)
	if err := r.Run(":8080"); err != nil {
		panic("gin 启动失败")
	}

}
