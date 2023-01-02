package model

func UpdateUserMsg(user User) error {
	err := DB.Save(&user).Error
	return err
}
func UpdatePost(post Post) error {
	err := DB.Save(&post).Error
	return err
}
