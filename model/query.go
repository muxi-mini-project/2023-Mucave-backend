package model

import "fmt"

func QueryNewPosts(start interface{}, length interface{}) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Select("*").Order("id desc").Offset(start).Limit(length).Find(&posts).Error
	return posts, err
}
func QueryHotPosts(start interface{}, length interface{}, ty string) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Where("type=?", ty).Select("*").Order("likes desc").Offset(start).Limit(length).Find(&posts).Error
	return posts, err
}
func QueryId(username string) (uint, error) {
	var user User
	err := DB.Model(&User{}).Select("id").Where("user_name=?", username).Find(&user).Error
	return user.ID, err
}
func QueryFollowing(UserId interface{}) ([]User, error) {
	var users []User
	err := DB.Table("relationships").Select("users.id,users.created_at,users.updated_at,users.deleted_at,users.name,users.gender,users.signature,users.avatar_path,users.birthday,users.hometown,users.grader,users.faculties").Joins("join users on relationships.followed_id = users.id ").Where("relationships.follower_id=?", UserId).Find(&users).Error
	fmt.Println(users)
	return users, err
}
func QueryFollowers(UserId interface{}) ([]User, error) {
	var users []User
	err := DB.Table("relationships").Select("users.id,users.created_at,users.updated_at,users.deleted_at,users.name,users.gender,users.signature,users.avatar_path,users.birthday,users.hometown,users.grader,users.faculties").Joins("join users on relationships.follower_id = users.id ").Where("relationships.followed_id=?", UserId).Find(&users).Error
	return users, err
}
func QueryUserPosts(users []User) ([]Post, error) {
	ids := make([]uint, 0)
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	var posts []Post
	err := DB.Where("author_id in (?)", ids).Find(&posts).Error
	return posts, err
}
func QueryIdPost(id string) (Post, error) {
	var post Post
	err := DB.Where("id = ?", id).Find(&post).Error
	return post, err
}
func QueryIdUser(id interface{}) (User, error) {
	var user User
	err := DB.Where("id = ?", id).Find(&user).Error
	return user, err
}
func QueryCommentByPostId(id string) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("post_id = ?", id).Order("created_at desc").Find(&comments).Error
	return comments, err
}
func QueryReplyByCommentId(id string) ([]Reply, error) {
	var replies []Reply
	err := DB.Where("comment_id = ?", id).Order("created_at desc").Find(&replies).Error
	return replies, err
}

type Outline struct {
	User        User
	PostNo      int
	FollowingNo int
	FollowerNo  int
}

func QueryOutline(id interface{}) Outline {
	var user User
	DB.Where("id = ?", id).Find(&user)
	var postNo, followingNo, followerNo int
	DB.Model(&Post{}).Where("author_id=?", id).Count(&postNo)
	DB.Model(&Relationship{}).Where("follower_id=?", id).Count(&followingNo)
	DB.Model(&Relationship{}).Where("followed_id=?", id).Count(&followerNo)
	outline := Outline{
		User:        user,
		PostNo:      postNo,
		FollowingNo: followingNo,
		FollowerNo:  followerNo,
	}
	return outline
}
func QueryPrivateMsg(sender interface{}, receiver interface{}) ([]PrivateMsg, error) {
	var privateMsgs []PrivateMsg
	err1 := DB.Model(&PrivateMsg{}).Where("sender_id=? AND receiver_id=? AND status=1", sender, receiver).Find(&privateMsgs).Error
	if err1 != nil {
		return privateMsgs, err1
	}
	ids := make([]uint, 0)
	for _, msg := range privateMsgs {
		ids = append(ids, msg.ID)
	}
	err2 := DB.Model(&PrivateMsg{}).Where("id in (?)", ids).Update("status", "0").Error
	if err2 != nil {
		return privateMsgs, err2
	}
	return privateMsgs, nil
}
func QueryMyComments(id interface{}) ([]Comment, error) {
	var comments []Comment
	err := DB.Model(&Comment{}).Where("user_id=?", id).Find(&comments).Error
	return comments, err
}
func QueryMyReplies(id interface{}) ([]Reply, error) {
	var replies []Reply
	err := DB.Model(&Reply{}).Where("from_who=?", id).Find(&replies).Error
	return replies, err
}
func QueryMyLikesPosts(id interface{}) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Omit("likes.id,likes.created_at,likes.updated_at,likes.deleted_at,likes.post_id,likes.user_id").Joins("join likes on likes.post_id=posts.id").Where("likes.user_id=?", id).Find(&posts).Error
	return posts, err
}