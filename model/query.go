package model

func QueryNewPosts(start int, length int) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Order("id desc").Offset(start).Limit(length).Find(&posts).Error
	return posts, err
}
func QueryHotPosts(startIndex int, length int, ty string, startTime string, endTime string) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Where("type=? && created_at between ? AND ?", ty, startTime, endTime).Order("likes desc").Offset(startIndex).Limit(length).Find(&posts).Error
	return posts, err
}
func QueryId(username string) (uint, error) {
	var user User
	err := DB.Model(&User{}).Select("id").Where("user_name=?", username).Take(&user).Error
	return user.ID, err
}
func QueryFollowing(UserId interface{}) ([]User, error) {
	var users []User
	err := DB.Table("relationships").Select("users.id,users.created_at,users.updated_at,users.deleted_at,users.name,users.gender,users.signature,users.avatar_path,users.birthday,users.hometown,users.grader,users.faculties").Joins("join users on relationships.followed_id = users.id ").Where("relationships.follower_id=?", UserId).Find(&users).Error
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
	err := DB.Where("author_id in (?)", ids).Order("id desc").Find(&posts).Error
	return posts, err
}
func QueryIdPost(id interface{}) (Post, error) {
	var post Post
	err := DB.Where("id = ?", id).Take(&post).Error
	return post, err
}
func QueryIdUser(id interface{}) (User, error) {
	var user User
	err := DB.Where("id = ?", id).Take(&user).Error
	return user, err
}
func QueryCommentByPostId(id string, my bool) ([]Comment, error) {
	var comments []Comment
	var err error
	if my {
		err = DB.Where("post_id = ?", id).Order("id desc").Find(&comments).Error
	} else {
		err = DB.Where("post_id = ? AND private = 0", id).Order("id desc").Find(&comments).Error
	}
	return comments, err
}

func QueryReplyByCommentId(commentId string, myId interface{}) ([]Reply, error) {
	var replies []Reply
	err := DB.Where("comment_id = ? AND private = 0 || object = ?", commentId, myId).Order("created_at desc").Find(&replies).Error
	return replies, err
}

type Outline struct {
	User        User
	PostNo      int64
	FollowingNo int64
	FollowerNo  int64
}

func QueryOutline(id interface{}) Outline {
	var user User
	DB.Where("id = ?", id).Find(&user)
	var postNo, followingNo, followerNo int64
	DB.Model(&Post{}).Where("author_id=?", id).Count(&postNo)
	DB.Model(&Relationship{}).Where("follower_id=?", id).Count(&followingNo)
	DB.Model(&Relationship{}).Where("followed_id=?", id).Count(&followerNo)
	outline := Outline{
		User:        user,
		PostNo:      postNo,      //动态数
		FollowingNo: followingNo, //关注数
		FollowerNo:  followerNo,  //动态数
	}
	return outline
}
func QueryPrivateMsg(sender interface{}, receiver interface{}) ([]PrivateMsg, error) {
	var privateMsgs []PrivateMsg
	var err1 error
	if s, _ := sender.(string); s == "all" {
		err1 = DB.Model(&PrivateMsg{}).Where("receiver_id=? AND status=1", receiver).Find(&privateMsgs).Error
	} else {
		err1 = DB.Model(&PrivateMsg{}).Where("sender_id=? AND receiver_id=? AND status=1", sender, receiver).Find(&privateMsgs).Error
	}
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
func WhetherLike(userId interface{}, postId interface{}) bool {
	var likes Likes
	err := DB.Where("user_id=? AND post_id=?", userId, postId).Take(&likes).Error
	if err != nil {
		return false
	}
	return true
}
func WhetherFollow(userId interface{}, myId interface{}) bool {
	var relationship Relationship
	err := DB.Where("followed_id = ? AND follower_id = ?", userId, myId).Take(&relationship).Error
	if err != nil {
		return false
	}
	return true
}
func SearchPosts(point interface{}) ([]Post, error) {
	var posts []Post
	p, _ := point.(string)
	err := DB.Where("title like ?", "%"+p+"%").Order("likes desc").Find(&posts).Error
	return posts, err
}
func QueryOneUserPosts(userId interface{}) ([]Post, error) {
	var posts []Post
	err := DB.Where("author_id=?", userId).Find(&posts).Error
	return posts, err
}

func QueryLikesSend(myPosts []Post) ([]Likes, error) {
	var postIds []uint
	for _, post := range myPosts {
		postIds = append(postIds, post.ID)
	}
	var likes []Likes
	err := DB.Model(&Likes{}).Where("post_id in (?) AND status = 1", postIds).Order("id desc").Find(&likes).Update("status", "0").Error
	return likes, err
}

func QueryCommentSend(myPosts []Post) ([]Comment, error) {
	var postIds []uint
	for _, post := range myPosts {
		postIds = append(postIds, post.ID)
	}
	var comments []Comment
	err := DB.Model(&Comment{}).Where("post_id in (?) AND status = 1 ", postIds).Order("id desc").Find(&comments).Update("status", "0").Error
	return comments, err
}

func QueryRepliesSend(myPosts []Post, myId interface{}) ([]Reply, error) {
	var postIds []uint
	for _, post := range myPosts {
		postIds = append(postIds, post.ID)
	}
	var replies1 []Reply
	err := DB.Model(&Reply{}).Where("post_id in (?) AND status_to_author = 1", postIds).Order("id desc").Find(&replies1).Update("status_to_author", "0").Error
	var replies2 []Reply
	err = DB.Model(&Reply{}).Where("object = ? AND status_to_object = 1", myId).Order("id desc").Find(&replies2).Update("status_to_object", "0").Error
	return append(replies1, replies2...), err
}
