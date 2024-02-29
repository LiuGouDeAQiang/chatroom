package tools

import (
	"github.com/bwmarrin/snowflake"
)

var snowNode *snowflake.Node

func GetUID() int64 {
	if snowNode == nil {
		snowNode, _ = snowflake.NewNode(1)
	}
	//node, _ := snowflake.NewNode(1)
	uid := snowNode.Generate().Int64()
	uid = uid % 100000000
	return uid
} //并发会出现重复，概率较低
