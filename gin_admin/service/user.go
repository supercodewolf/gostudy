package service

import (
	"gin_admin/dao"
	"gin_admin/model"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) error {
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &model.User{
		Username: username,
		Password: string(hashPwd),
	}
	return dao.CreateUser(user)
}

func Login(username, password string) (model.User, bool) {
	user, err := dao.FindUserByUsername(username)
	if err != nil {
		return user, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err == nil

}
