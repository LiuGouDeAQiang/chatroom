package model

import "fmt"

func GetGroupList(uid int64) ([]int64, error) {
	var groupList []UserGroups
	err := Conn.Where("user_uid=?", uid).Find(&groupList).Error
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	var groupUIDs []int64
	for _, userGroup := range groupList {
		groupUIDs = append(groupUIDs, userGroup.GroupUid)
	}
	return groupUIDs, err
}
func GetFriendsList(uid int64) ([]int64, error) {
	var FriendsList []UserFriends
	err := Conn.Where("user_uid=?", uid).Find(&FriendsList).Error
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	var groupUIDs []int64
	for _, userFriends := range FriendsList {
		groupUIDs = append(groupUIDs, userFriends.FriendUid)
	}
	return groupUIDs, err
}
func GetRenList(uid int64) ([]int64, error) {
	var groupList []UserGroups
	err := Conn.Where("user_uid=?", uid).Find(&groupList).Error
	if err != nil {
		fmt.Println("获取群列表失败")
	}
	var groupUIDs []int64
	for _, userGroup := range groupList {
		groupUIDs = append(groupUIDs, userGroup.GroupUid)
	}
	return groupUIDs, err
}
