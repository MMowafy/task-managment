package models

import (
	"time"
)

type Notification struct {
	Id        int       `gorm:"primary_key;column:id;" json:"id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	TaskId    int       `gorm:"column:task_id;index:task_id" json:"task_id"`
	IsRead    bool      `gorm:"column:is_read;index:is_read" json:"is_read"`
	Message   string    `gorm:"column:message" json:"message"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func NewNotification() *Notification {
	return &Notification{}
}
