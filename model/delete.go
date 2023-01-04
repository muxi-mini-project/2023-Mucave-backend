package model

import (
	"github.com/jinzhu/gorm"
)

func DeleteRelationship(followedId interface{}, followerId interface{}) error {
	err := DB.Where("followed_id=? AND follower_id= ?", followedId, followerId).Delete(&Relationship{}).Error
	return err
}
func DeletePost(postId interface{}) error {
	err := DB.Where("id=?", postId).Delete(&Post{}).Error
	return err
}

func DeleteLikes(userId interface{}, postId interface{}) error {
	err := DB.Model(&Post{}).Where("id=?", postId).Update("likes", gorm.Expr("likes-1")).Error
	if err != nil {
		return err
	}
	err = DB.Delete(&Likes{}).Where("post_id=? AND author_id=?", postId, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId interface{}) error {
	err := DB.Where("id=?", commentId).Delete(&Comment{}).Error
	return err
}

func DeleteReply(replyId interface{}) error {
	err := DB.Where("id=?", replyId).Delete(&Reply{}).Error
	return err
}
