package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
	"strings"
	"task-managment/api/utils"
	"time"
)

type User struct {
	Id        int       `gorm:"primary_key;column:id;" json:"id"`
	Email     string    `gorm:"column:email;index:email" json:"email"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Tasks     []Task    `json:"tasks" gorm:"PRELOAD:false"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Format() *User {
	user.Email = strings.ToLower(user.Email)
	return user
}

func (user *User) Sanitize() *User {
	user.Email = utils.SanitizeString(user.Email)
	user.FirstName = utils.SanitizeString(user.FirstName)
	user.LastName = utils.SanitizeString(user.LastName)
	return user
}

func (user *User) Validate() error {
	err := user.validateStruct()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) validateStruct() error {

	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.FirstName, validation.Required, validation.Length(1, 25), validation.Match(regexp.MustCompile("^[a-zA-Z ]+$"))),
		validation.Field(&user.LastName, validation.Required, validation.Length(1, 25), validation.Match(regexp.MustCompile("^[a-zA-Z ]+$"))),
	)
}
