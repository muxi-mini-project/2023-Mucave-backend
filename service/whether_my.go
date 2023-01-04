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

func WhetherMyComment(userId interface{}, commentId interface{}) bool {
	comments, _ := model.QueryMyComments(userId)
	for _, comment := range comments {
		if strconv.Itoa(int(comment.ID)) == commentId {
			return true
		}
	}
	return false

}

func WhetherMyReply(userId interface{}, replyId interface{}) bool {
	replies, _ := model.QueryMyReplies(userId)
	for _, reply := range replies {
		if strconv.Itoa(int(reply.ID)) == replyId {
			return true
		}
	}
	return false
}
