package model

func UpdateUserMsg(user User, userId interface{}) error {
	err := DB.Model(&User{}).Where("id=?", userId).Updates(user).Error
	//User{Name: user.Name, Gender: user.Gender, Signature: user.Signature, Birthday: user.Birthday, Hometown: user.Hometown, Grader: user.Grader, Faculties: user.Faculties}
	return err
}
func UpdatePost(file string, post Post, posId interface{}) error {
	var err error
	if file == "yes" {
		err = DB.Model(&Post{}).Where("id = ?", posId).Updates(post).Error
	} else {
		err = DB.Model(&Post{}).Where("id = ?", posId).Updates(post).Error
		//Post{Title: post.Title, AuthorId: post.AuthorId, Type: post.Type, Content: post.Content}
	}
	return err
}
func UpdateAvatar(id interface{}, avatarPath interface{}) error {
	err := DB.Model(&User{}).Where("id = ?", id).Update("avatar_path", avatarPath).Error
	return err
}
