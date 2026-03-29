package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string `gorm:"type:varchar(200);not null;index" json:"title"`
	Summary   string `gorm:"type:varchar(500)" json:"summary,omitempty"`
	Content   string `gorm:"type:text;not null" json:"content"`
	ViewCount uint   `gorm:"default:0" json:"view_count"`
	AuthorID  uint   `gorm:"not null;index" json:"author_id"`
	Tags      []Tag  `gorm:"many2many:article_tags;" json:"tags,omitempty"`
}
