package logic

import (
	model "chatroom/app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	UID     string // 用户 ID
	Item    string // 聊天项
	Content string // 消息内容
}

var broadcast = make(chan string)
var chatRooms = make(map[string]map[string]map[*websocket.Conn]struct{})

// var chatRoomsMutex sync.RWMutex
var item string
var uidStr string

// 读取ws信息
// 处理 WebSocket 消息

func HandWebSocket(c *gin.Context) {
	item = c.Query("item")
	uidStr = c.Query("uid")
	fmt.Println("item和uid分别是", item, uidStr)
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		// 处理转换错误
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
		return
	}
	defer conn.Close()
	//chatRoomsMutex.Lock()
	//defer chatRoomsMutex.Unlock()
	// 创建聊天室的连接 map
	if _, ok := chatRooms[item]; !ok {
		chatRooms[item] = make(map[string]map[*websocket.Conn]struct{})

	}
	if _, ok := chatRooms[item][uidStr]; !ok {
		chatRooms[item][uidStr] = make(map[*websocket.Conn]struct{})
	}
	// 将连接添加到特定聊天室
	chatRooms[item][uidStr][conn] = struct{}{}
	//发送历史消息
	err = SendChatHistory(c, conn, item)
	if err != nil {
		fmt.Println("发送历史消息失败")
	}
	for {
		// 读取 WebSocket 消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			delete(chatRooms[item][uidStr], conn) //   disconnected client from the map
			break
		}
		Msg := string(msg)
		// 处理消息
		fmt.Printf("Received message from %v: %s\n", uid, Msg)
		//将聊天信息添加到mysql以及redis缓存中
		blockList, err := model.Block()
		if err != nil {
			fmt.Println("敏感词获取失败")
		}
		// 遍历消息进行屏蔽词替换
		for _, blockedWord := range blockList {
			Msg = strings.ReplaceAll(Msg, blockedWord, "*")
		}
		history := model.MesHistory{
			UserUid:     int64(uid),
			Content:     Msg,
			CreatedTime: time.Now(),
		}
		err = model.AddRedisHistory(c, item, Msg)
		if err != nil {
			fmt.Println("聊天信息缓存失败")
		}
		err = model.AddHistoryToTable(item, history)
		if err != nil {
			fmt.Println("聊天信息存储失败", err)
		}
		// 将消息发送给其他连接的客户端
		broadcast <- Msg
	}
}

// HandleMessages 发送接收到的消息
func HandleMessages() {
	for {
		msg := <-broadcast
		fmt.Println(msg)
		//chatRoomsMutex.RLock()
		// 将消息发送给特定聊天室的所有连接
		if connections, ok := chatRooms[item][uidStr]; ok {
			fmt.Println("item和uidStr", item, uidStr)
			for connection := range connections {
				err := connection.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("Write error: ", err)
					fmt.Println("发送消息出现错误")
					connection.Close()
					delete(chatRooms[item][uidStr], connection)
				}
			}
		}
		//chatRoomsMutex.RUnlock()
	}
}

// SendChatHistory 发送历史信息
func SendChatHistory(c *gin.Context, client *websocket.Conn, item string) error {
	// 从 Redis 获取聊天历史消息
	chatHistory, err := model.GetRedisHistory(c, item)
	if err != nil {
		fmt.Println("缓存不存在")
		chatHistory, err = model.GetHistoryFromTable(item)
		if err != nil {
			return errors.New("创建表的时候错误")
		}
	}
	fmt.Printf("聊天历史%s", chatHistory)
	// 将聊天历史消息发送给当前客户端
	for _, msg := range chatHistory {
		fmt.Println("聊天历史是", msg)
		err := client.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Write error: ", err)
			client.Close()
			return err
		}
	}

	return nil
}
