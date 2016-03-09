package models

import (
	"time"

	"modules/log"
)

func (m model) GetPostById(id uint64) *Post {
	post := Post{}
	if err := m.db.Where("id = ?", id).First(&post).Error; err != nil {
		log.DebugPrint("Get post error: %v", err)
		return nil
	}

	if err := m.db.Model(&post).Related(&post.User).Error; err != nil {
		log.DebugPrint("Post user related error: %v", err)
		return &post
	}

	return &post
}

func (m model) GetUserPostsByUserId(userId uint64, page int, size int) *[]Post {
	posts := []Post{}
	if err := m.db.Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
		log.DebugPrint("Get user posts error: %v", err)
		return nil
	}

	for key, post := range posts {
		if err := m.db.Model(&post).Related(&post.User).Error; err != nil {
			log.DebugPrint("Post user related error: %v", err)
		}
		posts[key] = post
	}

	return &posts
}

func (m model) PostSave() {
	tx := m.db.Begin()

	post1 := Post{Title: "标题3"}
	if err := tx.Create(&post1).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	post2 := Post{Title: "标题4"}
	if err := tx.Create(&post2).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}

type Post struct {
	Id        uint64    `json:"id,omitempty"`
	UserId    uint64    `form:"user_id" json:"user_id,omitempty"`
	Title     string    `form:"title" json:"title,omitempty"`
	Context   string    `form:"context" json:"context,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	User User `gorm:"ForeignKey:UserId;AssociationForeignKey:Id" json:"user"`
}

func (p Post) TableName() string {
	return "post"
}
