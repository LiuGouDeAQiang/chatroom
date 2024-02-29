package app

import (
	"chatroom/app/models"
	"chatroom/app/router"
)

func Start() {
	model.NewMysql()
	model.NewRdb()
	defer func() {
		model.Close()
	}()
	router.New()

}
