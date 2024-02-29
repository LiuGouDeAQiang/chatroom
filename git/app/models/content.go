package model

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateHistoryTableIfNotExists 根据 item 创建历史记录表（如果不存在）
func CreateHistoryTableIfNotExists(item string) error {
	tableName := fmt.Sprintf("%sMesHistory", item)
	// 检查表是否存在
	if !Conn.Migrator().HasTable(tableName) {
		// 如果不存在，则创建表
		if err := Conn.Table(tableName).AutoMigrate(MesHistory{}); err != nil {
			return errors.New("创建表历史失败")
		}
	}
	return nil
}

// AddHistoryToTable 添加历史记录到指定表
func AddHistoryToTable(item string, history MesHistory) error {
	tableName := fmt.Sprintf("%sMesHistory", item)
	err := CreateHistoryTableIfNotExists(item)
	if err != nil {
		return errors.New("在进行添加的时候创建表单失败")
	}
	// 插入历史记录
	if err := Conn.Table(tableName).Create(&history).Error; err != nil {
		return err
	}

	return nil
}

// GetHistoryFromTable 获取指定表的历史记录
func GetHistoryFromTable(item string) ([]string, error) {
	err := CreateHistoryTableIfNotExists(item)
	if err != nil {
		return nil, errors.New("创建表单失败")
	}
	tableName := fmt.Sprintf("%sMesHistory", item)
	var history []MesHistory
	if err := Conn.Table(tableName).Find(&history).Error; err != nil {
		return nil, err
	}
	var contentList []string

	for _, entry := range history {
		contentList = append(contentList, entry.Content)
	}
	return contentList, nil
}

func AddRedisHistory(c *gin.Context, item string, message string) error {
	// 存储 JSON 字符串到 Redis
	err := Rdb.LPush(c, "chat_messages_"+item, message).Err()
	if err != nil {
		return err
	}

	return nil

}

func GetRedisHistory(c *gin.Context, item string) ([]string, error) {
	// 获取 Redis 中的历史记录
	result, err := Rdb.LRange(c, "chat_messages_"+item, 0, 4).Result()
	if err != nil {
		return nil, err
	}

	// 颠倒顺序
	reversedResult := make([]string, len(result))
	for i, j := 0, len(result)-1; i < len(result); i, j = i+1, j-1 {
		reversedResult[i] = result[j]
	}

	return reversedResult, nil
}
