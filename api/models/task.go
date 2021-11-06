package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"task-managment/api/utils"
	"time"
)

type Task struct {
	Id        int       `gorm:"primary_key;column:id;" json:"id"`
	UserId    int       `gorm:"column:user_id;index:user_id" json:"user_id"`
	Name      string    `gorm:"column:name;index:name" json:"name"`
	Duration  int       `gorm:"column:duration;index:duration" json:"duration"`
	StartAt   time.Time `gorm:"column:start_at" json:"start_at"`
	EndAt     time.Time `gorm:"column:end_at" json:"end_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func NewTask() *Task {
	return &Task{}
}

func (task *Task) Format() *Task {
	return task
}

func (task *Task) Sanitize() *Task {
	task.Name = utils.SanitizeString(task.Name)
	return task
}

func (task *Task) Validate() error {
	err := task.validateStruct()
	if err != nil {
		return err
	}
	return nil
}

func (task *Task) validateStruct() error {

	return validation.ValidateStruct(task,
		validation.Field(&task.Name, validation.Required),
		validation.Field(&task.UserId, validation.Required),
		validation.Field(&task.Duration, validation.Required),
		validation.Field(&task.StartAt, validation.Required),
	)
}
