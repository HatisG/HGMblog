package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Posts []Article `gorm:"many2many:article_tags;" json:"posts,omitempty"`
}
