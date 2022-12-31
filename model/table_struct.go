package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName   string
	Name       string //名字
	Gender     string
	Signature  string //个性签名
	AvatarPath string //头像
	Birthday   time.Time
	Hometown   string
	Grader     string
	Faculties  string //院系
}
type Post struct {
	gorm.Model
	AuthorId  uint
	Type      string
	Content   string
	Likes     int
	CommentNo int
	FilePath1 string
	FilePath2 string
	FilePath3 string
	FilePath4 string
}
type Comment struct {
	gorm.Model
	PostId  uint
	UserId  uint
	Content string
}
type Reply struct {
	gorm.Model
	FromWho   uint
	CommentId uint
	Object    uint
	Content   string
}
type Relationship struct {
	gorm.Model
	FollowerId uint
	FollowedId uint
}
type PrivateMsg struct {
	gorm.Model
	SenderId   int
	ReceiverId int
	Status     int
	Content    string
	FilePath   string
	SendTime   time.Time
}
type Likes struct {
	gorm.Model
	PostId uint
	UserId uint
}
