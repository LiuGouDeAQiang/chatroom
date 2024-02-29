package model

import "time"

type User struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Uid         int64     `gorm:"column:uid;type:bigint(20);NOT NULL" json:"uid"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Password    string    `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	Telephone   string    `gorm:"column:telephone;type:varchar(255);NOT NULL" json:"telephone"`
	Email       string    `gorm:"column:email;type:varchar(255)" json:"email"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime" json:"updated_time"`
}

func (m *User) TableName() string {
	return "user"
}

type Group struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Uid         int64     `gorm:"column:uid;NOT NULL"`
	Name        string    `gorm:"column:name;NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL"`
}

func (g *Group) TableName() string {
	return "group"
}

type UserFriends struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserUid     int64     `gorm:"column:user_uid;NOT NULL"`
	FriendUid   int64     `gorm:"column:friend_uid;NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL"`
}

func (u *UserFriends) TableName() string {
	return "user_friends"
}

type UserGroups struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserUid     int64     `gorm:"column:user_uid;NOT NULL"`
	GroupUid    int64     `gorm:"column:group_uid;NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL"`
}

func (u *UserGroups) TableName() string {
	return "user_groups"
}

type MesHistory struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserUid     int64     `gorm:"column:user_uid;default:NULL"`
	Content     string    `gorm:"column:content;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	DelTime     time.Time `gorm:"column:del_time;default:NULL"`
}

func (m *MesHistory) TableName() string {
	return "mes_history"
}

type BlockingWords struct {
	Id      int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Content string `gorm:"column:content;default:NULL"`
}

func (b *BlockingWords) TableName() string {
	return "blocking_words"
}
