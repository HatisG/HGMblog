package service

import (
	"errors"

	"HGMblog_v1.0/dao"
	"HGMblog_v1.0/model"
)

type ArticleService struct {
	ArticleDao *dao.ArticleDao
	TagDao     *dao.TagDao
}

type CreateRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Summary string   `json:"summary,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type UpdateRequest struct {
	Title   *string  `json:"title"`
	Content *string  `json:"content"`
	Summary *string  `json:"summary,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

func (s *ArticleService) Create(req *CreateRequest, authorID uint) error {

	tags, err := s.TagDao.FindOrCreateBatch(req.Tags)
	if err != nil {
		return err
	}

	article := &model.Article{
		Title:    req.Title,
		Content:  req.Content,
		Summary:  req.Summary,
		AuthorID: authorID,
		Tags:     tags,
	}
	return s.ArticleDao.Create(article)
}

func (s *ArticleService) Update(id uint, req *UpdateRequest, userID uint) error {
	article, err := s.ArticleDao.FindByID(id)
	if err != nil {
		return err
	}

	if article.AuthorID != userID {
		return errors.New("无权限修改")
	}

	if req.Title != nil {
		article.Title = *req.Title
	}
	if req.Content != nil {
		article.Content = *req.Content
	}
	if req.Summary != nil {
		article.Summary = *req.Summary
	}

	if req.Tags != nil {
		tags, err := s.TagDao.FindOrCreateBatch(req.Tags)
		if err != nil {
			return err
		}
		article.Tags = tags
	}

	return s.ArticleDao.Update(article)
}

func (s *ArticleService) SearchPublic(keyword, tag string, page, pagesize int) ([]model.Article, int64, error) {
	query := &dao.ArticleQuery{
		Keyword:  keyword,
		Tag:      tag,
		Page:     page,
		Pagesize: pagesize,
	}

	return s.ArticleDao.List(query)

}

func (s *ArticleService) SearchByAuthor(userID uint, page, pagesize int) ([]model.Article, int64, error) {
	query := &dao.ArticleQuery{
		AuthorID: userID,
		Page:     page,
		Pagesize: pagesize,
	}

	return s.ArticleDao.List(query)

}

func (s *ArticleService) Delete(userID uint, articleID uint) error {
	article, err := s.ArticleDao.FindByID(articleID)
	if err != nil {
		return errors.New("文章不存在")
	}
	if article.AuthorID != userID {
		return errors.New("无权限删除")
	}
	return s.ArticleDao.Delete(articleID)
}

func (s *ArticleService) Get(id uint) (*model.Article, error) {
	return s.ArticleDao.FindByID(id)
}
