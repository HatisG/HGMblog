package dao

import (
	"HGMblog_v1.0/model"
	"gorm.io/gorm"
)

type ArticleDao struct {
	DB *gorm.DB
}

type ArticleQuery struct {
	Keyword  string
	Tag      string
	AuthorID uint
	Page     int
	Pagesize int
	SortBy   string
	Order    string
}

func (dao *ArticleDao) Create(article *model.Article) error {
	return dao.DB.Create(article).Error
}

func (dao *ArticleDao) Update(article *model.Article) error {
	return dao.DB.Save(article).Error
}

func (dao *ArticleDao) List(query *ArticleQuery) ([]model.Article, int64, error) {
	//构建查询
	db := dao.DB.Model(&model.Article{})

	//模糊查询
	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		db = db.Where("title like ? or content like ?", keyword, keyword)
	}

	if query.Tag != "" {
		db = db.Joins("join article_tags on articles.id = article_tags.article_id").
			Joins("join tags on tags.id = article_tags.tag_id").
			Where("tags.name = ?", query.Tag)
	}

	if query.AuthorID != 0 {
		db = db.Where("author_id = ?", query.AuthorID)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if query.SortBy != "" {
		order := query.SortBy + " " + query.Order
		db = db.Order(order)
	} else {
		db = db.Order("created_at DESC")
	}

	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Pagesize <= 0 {
		query.Pagesize = 10
	}

	offset := (query.Page - 1) * query.Pagesize
	db = db.Offset(offset).Limit(query.Pagesize)

	var articles []model.Article
	err := db.Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil

}

func (dao *ArticleDao) Delete(id uint) error {
	return dao.DB.Delete(&model.Article{}, id).Error
}

func (dao *ArticleDao) FindByID(id uint) (*model.Article, error) {
	var article model.Article
	err := dao.DB.Preload("Tags").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}
