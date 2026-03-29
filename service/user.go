package service

import (
	"errors"

	"HGMblog_v1.0/dao"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserDao     *dao.UserDao
	AuthService *AuthService
}

func (s *UserService) Register(username, password string) error {
	//查找username
	_, err := s.UserDao.SearchByUsername(username)
	//存在username
	if err == nil {
		return errors.New("用户名已存在")
	}
	//数据库错误
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//加密
	Hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//创建一个user
	_, err = s.UserDao.Create(username, string(Hashpassword), username)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(username, password string) (string, error) {
	//username查找
	user, err := s.UserDao.SearchByUsername(username)
	//用户名或密码错误
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("用户名或密码错误")
	}

	token, err := s.AuthService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *UserService) Delete(userID uint, username string) error {
	//用户注销账号
	user, err := s.UserDao.SearchByUsername(username)
	if err != nil {
		return errors.New("无权限注销用户")
	}
	if user.ID == userID {
		return s.UserDao.Delete(username)
	} else {
		return errors.New("无权限注销用户")
	}
}
