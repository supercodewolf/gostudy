package dao

import (
	"gin_admin/config"
	"gin_admin/model"
)

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error

}

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := config.DB.Where("username =?", username).First(&user).Error
	return user, err
}
func FindAllUsers() ([]model.User, error) {
	var users []model.User
	err := config.DB.Find(&users).Error
	return users, err
}

func FindUserByID(id uint) (model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return user, err
}
func UpdateUsername(id uint, username string) error {
	return config.DB.Model(&model.User{}).Where("id=?", id).Update("username", username).Error
}

func DeleteUser(id uint) error {
	// return config.DB.Delete(&model.User{}, id).Error
	// 第一步：先查用户是否存在
	var user model.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		return err // 用户不存在，返回错误
	}

	// 第二步：用户存在，执行删除
	return config.DB.Delete(&user).Error

}
