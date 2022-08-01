package domain

import (
	"time"
)

type Users struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Username  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Referrer  string `json:"referrer"`
}

type Content struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `json:"title"`
	Title_content string `json:"title_content"`
	Content       string `json:"content"`
	Username      string `json:"username"`
}

type DeleteContent struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `json:"title"`
	Title_content string `json:"title_content"`
	Username      string `json:"username"`
}

type News struct {
	ID             int       `gorm:"primaryKey;serializer" json:"id"`
	CreatedAt      time.Time `json:"created_at" binding:"required"`
	Title          string    `json:"title" binding:"required"`
	TitleContent   string    `json:"title_content" binding:"required"`
	Content        string    `json:"content" binding:"required"`
	CreatorContent string    `gorm:"foreignKey:UserName" binding:"required" json:"creator_content"`
	Likes          int       `gorm:"default:0" json:"likes"`
}

type DotEnv struct {
	Host string
	User string
	PW   string
	Name string
	Port string
	SSL  string
}
