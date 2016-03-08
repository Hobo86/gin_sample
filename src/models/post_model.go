package models

import (
	"time"

	"modules/log"
)

func (m model) GetPostById(id int64) *Post {
	post := Post{}
	if err := m.db.Where("id = ?", id).First(&post).Related(&post.User).Error; err != nil {
		return nil
	}

	return &post
}

func (m model) GetUserPostsByUserId(userId int64, page int, size int) *[]Post {
	posts := []Post{}
	if err := m.db.Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
		return nil
	}

	for key, post := range posts {
		if err := m.db.Model(&post).Related(&post.User).Error; err != nil {
			log.DebugPrint("User Related Error: %v", err)
		}
		posts[key] = post
	}

	return &posts
}

type Post struct {
	Id        int64
	UserId    int64     `form:"user_id"`
	Title     string    `form:"title"`
	CreatedAt time.Time `gorm:"column:created_time"`
	UpdatedAt time.Time `gorm:"column:updated_time"`

	User User
}

func (p Post) TableName() string {
	return "post"
}
