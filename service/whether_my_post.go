package service

import (
	"Mucave/model"
	"strconv"
)

func WhetherMyPost(userId interface{}, postId interface{}) bool {
	posts, _ := model.QueryOneUserPosts(userId)
	for _, post := range posts {
		if strconv.Itoa(int(post.ID)) == postId {
			return true
		}
	}
	return false
}
