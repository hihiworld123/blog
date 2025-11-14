package entity

import "time"

type User struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;size:50;not null" json:"username"`
	Password string `gorm:"not null;size:255"`
	Email    string `gorm:"unique;size:50;not null" json:"email"`
	Posts    []Post `gorm:"foreignKey:UserId"`
}

type Post struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null;size:255" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	UserId    int64     `gorm:"not null;index" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	PostId    int64     `gorm:"not null;index" json:"post_id"`
	UserId    int64     `gorm:"not null;index" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
