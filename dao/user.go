package dao

import (
	"HGMblog_v1.0/model"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func (dao *UserDao) Create(username, password, nickname string) (*model.User, error) {
	//创建user结构体
	user := &model.User{
		UserName: username,
		Password: password,
		NickName: nickname,
	}
	//插入数据库
	err := dao.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (dao *UserDao) SearchByUsername(username string) (*model.User, error) {
	var user model.User

	//数据库根据username查找
	err := dao.DB.Where("user_name = ?", username).First(&user).Error

	//不同返回值有不同的用途
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (dao *UserDao) SearchByNickname(nickname string) ([]model.User, error) {
	var users []model.User
	//数据库根据nickname查找
	err := dao.DB.Where("nick_name = ?", nickname).Find(&users).Error

	//不同返回值有不同的用途
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (dao *UserDao) Delete(username string) error {
	return dao.DB.Where("user_name = ?", username).Delete(&model.User{}).Error
}
