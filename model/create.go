package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func CreateUser(username string) User {
	var user1 User
	user := User{
		UserName: username,
		Birthday: time.Now(),
	}
	DB.Create(&user).Find(&user1)
	return user1
}
func CreatePost(post Post) error {
	err := DB.Create(&post).Error
	return err
}
func CreateReply(reply Reply) error {
	err := DB.Create(&reply).Error
	return err
}
func CreateComment(comment Comment) error {
	err := DB.Model(&Post{}).Where("id=?", comment.PostId).Update("comment_no", gorm.Expr("comment_no+1")).Error
	if err != nil {
		return err
	}
	err = DB.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateLikes(likes Likes) error {
	err := DB.Model(&Post{}).Where("id=?", likes.PostId).Update("likes", gorm.Expr("likes+1")).Error
	if err != nil {
		return err
	}
	err = DB.Create(&likes).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateRelationship(relationship Relationship) error {
	err := DB.Create(&relationship).Error
	return err
}
func CreatePrivateMsg(privateMsg PrivateMsg) error {
	err := DB.Create(&privateMsg).Error
	return err
}
