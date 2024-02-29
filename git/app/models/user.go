package model

import "fmt"

func GetUser(uid int64) *User {
	var ret User
	if err := Conn.Table("user").Where("uid=?", uid).Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}
func CreateUser(user *User) error {
	fmt.Println(user)
	if err := Conn.Create(user).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
