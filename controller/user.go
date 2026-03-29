package controller

import (
	"net/http"

	"HGMblog_v1.0/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func (ctrl *UserController) Register(c *gin.Context) {
	var user struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := ctrl.UserService.Register(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})

}

func (ctrl *UserController) Login(c *gin.Context) {
	var user struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ctrl.UserService.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

func (ctrl *UserController) Delete(c *gin.Context) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	username := c.Param("username")
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID类型错误"})
		return
	}
	err := ctrl.UserService.Delete(userID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注销成功"})
}
