package controller

import (
	"net/http"
	"strconv"

	"HGMblog_v1.0/service"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	ArticleService *service.ArticleService
}

func (ctrl *ArticleController) Create(c *gin.Context) {
	var req service.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idAny, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	userID, ok := idAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID类型错误"})
		return
	}
	err := ctrl.ArticleService.Create(&req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func (ctrl *ArticleController) Update(c *gin.Context) {
	var req service.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	temp, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效文章id"})
		return
	}
	articleID := uint(temp)

	userIDAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID类型错误"})
		return
	}

	err = ctrl.ArticleService.Update(articleID, &req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "修改成功"})

}

func (ctrl *ArticleController) SearchByAuthor(c *gin.Context) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID类型错误"})
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pagesize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	articles, total, err := ctrl.ArticleService.SearchByAuthor(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pagesize": pageSize,
	})

}

func (ctrl *ArticleController) SearchPublic(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	tag := c.DefaultQuery("tag", "")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pagesize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	articles, total, err := ctrl.ArticleService.SearchPublic(keyword, tag, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pagesize": pageSize,
	})

}

func (ctrl *ArticleController) Delete(c *gin.Context) {
	userIDAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID类型错误"})
		return
	}

	idParam := c.Param("id")
	temp, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效文章id"})
		return
	}
	articleID := uint(temp)

	err = ctrl.ArticleService.Delete(userID, articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})

}

func (ctrl *ArticleController) Get(c *gin.Context) {
	idParam := c.Param("id")
	temp, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效文章id"})
		return
	}
	articleID := uint(temp)

	article, err := ctrl.ArticleService.Get(articleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)

}
