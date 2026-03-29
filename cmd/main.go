package main

import (
	"fmt"
	"log"

	"HGMblog_v1.0/config"
	"HGMblog_v1.0/controller"
	"HGMblog_v1.0/dao"
	"HGMblog_v1.0/middleware"
	"HGMblog_v1.0/model"
	"HGMblog_v1.0/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//加载配置
	config.Load()

	//读取配置
	dsn := config.Get("DSN", "")
	if dsn == "" {
		log.Fatal("DSN 未配置")
	}
	jwtSecret := config.Get("JWT_SECRET", "")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET 未配置")
	}
	port := config.Get("PORT", "")
	if port == "" {
		log.Fatal("PORT 未配置")
	}

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	//自动迁移
	if err := db.AutoMigrate(&model.User{}, &model.Article{}, &model.Tag{}); err != nil {
		log.Fatal("数据库迁移失败", err)
	}
	fmt.Println("数据库连接成功")

	//初始化依赖
	userDao := &dao.UserDao{DB: db}
	articleDao := &dao.ArticleDao{DB: db}
	tagDao := &dao.TagDao{DB: db}
	authService := &service.AuthService{SecretKey: []byte(jwtSecret)}
	userService := &service.UserService{
		UserDao:     userDao,
		AuthService: authService,
	}
	articleService := &service.ArticleService{
		ArticleDao: articleDao,
		TagDao:     tagDao,
	}
	userController := &controller.UserController{
		UserService: userService,
	}
	articleController := &controller.ArticleController{
		ArticleService: articleService,
	}

	//gin初始化
	r := gin.Default()

	//public路由
	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
		public.GET("/articles/:id", articleController.Get)
		public.GET("/articles", articleController.SearchPublic)
	}

	//认证路由
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(authService))
	{
		auth.DELETE("/user/:username", userController.Delete)
		auth.POST("/articles", articleController.Create)
		auth.PUT("/articles/:id", articleController.Update)
		auth.DELETE("/articles/:id", articleController.Delete)
		auth.GET("/my/articles", articleController.SearchByAuthor)
	}

	//启动服务
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("服务启动失败", err)
	} else {
		fmt.Println("服务启动成功")
	}

}
