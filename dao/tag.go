package dao

import (
	"HGMblog_v1.0/model"
	"gorm.io/gorm"
)

type TagDao struct {
	DB *gorm.DB
}

func (dao *TagDao) FindOrCreateBatch(names []string) ([]model.Tag, error) {
	var tags []model.Tag
	for _, name := range names {
		var tag model.Tag
		err := dao.DB.Where("name = ?", name).FirstOrCreate(&tag, model.Tag{Name: name}).Error
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil

}
