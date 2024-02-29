package model

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	NewMysql()
	block, err := Block()
	if err != nil {
		return
	}
	fmt.Println(block)
}
