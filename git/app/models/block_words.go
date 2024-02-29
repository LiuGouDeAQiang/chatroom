package model

import "fmt"

func Block() ([]string, error) {
	var blockList []BlockingWords
	if err := Conn.Find(&blockList).Error; err != nil {
		fmt.Printf("屏蔽词获取失败 err:%s", err.Error())
		return nil, err
	}

	var contentList []string
	result := Conn.Model(&BlockingWords{}).Pluck("content", &contentList)
	if result.Error != nil {
		fmt.Printf("内容获取失败 err:%s", result.Error.Error())
		return nil, result.Error
	}
	return contentList, nil

}
