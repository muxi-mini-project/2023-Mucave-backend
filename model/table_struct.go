package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName   string `form:"user_name" json:"user_name"`
	Name       string `form:"name" json:"name"`
	Gender     string `form:"gender" json:"gender"`
	Signature  string `form:"signature" json:"signature"`
	AvatarPath string `form:"avatar_path" json:"avatar_path"`
	Birthday   time.Time
	Hometown   string `form:"hometown" json:"hometown"`
	Grader     string `form:"grader" json:"grader"`
	Faculties  string `form:"faculties" json:"faculties"`
}
type Post struct {
	gorm.Model
	Title     string `form:"title" json:"title"`
	AuthorId  uint   `form:"author_id" json:"author_id"`
	Type      string `form:"type" json:"type"`
	Content   string `form:"content" json:"content"`
	Likes     int    `form:"likes" json:"likes"`
	CommentNo int    `form:"comment_no" json:"comment_no"`
	FilePath1 string `form:"file_path1" json:"file_path1"`
	FilePath2 string `form:"file_path2" json:"file_path2"`
	FilePath3 string `form:"file_path3" json:"file_path3"`
	FilePath4 string `form:"file_path4" json:"file_path4"`
	FilePath5 string `form:"file_path5" json:"file_path5"`
	FilePath6 string `form:"file_path6" json:"file_path6"`
}
type Comment struct {
	gorm.Model
	PostId  uint   `form:"post_id" json:"post_id"`
	UserId  uint   `json:"user_id" form:"user_id"`
	Content string `json:"content" form:"content"`
}
type Reply struct {
	gorm.Model
	FromWho   uint   `json:"from_who" form:"from_who"`
	CommentId uint   `json:"comment_id" form:"comment_id"`
	Object    uint   `json:"object" form:"object"`
	Content   string `json:"content" form:"content"`
}
type Relationship struct {
	gorm.Model
	FollowerId uint `json:"follower_id" form:"follower_id"`
	FollowedId uint `json:"followed_id" form:"followed_id"`
}
type PrivateMsg struct {
	gorm.Model
	SenderId   uint      `json:"sender_id" form:"sender_id"`
	ReceiverId uint      `json:"receiver_id" form:"receiver_id"`
	Status     int       `json:"status" form:"status"`
	Content    string    `json:"content" form:"content"`
	FilePath   string    `json:"file_path" form:"file_path"`
	SendTime   time.Time `json:"send_time" form:"send_time"`
}
type Likes struct {
	gorm.Model
	PostId uint `json:"post_id" form:"post_id"`
	UserId uint `json:"user_id" form:"user_id"`
}
