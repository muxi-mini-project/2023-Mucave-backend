package model

func UpdateUserMsg(user User) error {
	err := DB.Save(&user).Error
	return err
}
