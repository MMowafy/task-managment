package infrastructure

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"task-managment/api/application"
	"task-managment/api/models"
	"task-managment/api/utils"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db, _ := application.GetPostgresConnectionByName("appdb")
	userRepository := &UserRepository{
		db,
	}
	userRepository.db.LogMode(application.GetConfig().GetBool("app.db_logs"))
	return userRepository
}

func (userRepository *UserRepository) checkDBConnection() error {
	if userRepository.db == nil {
		application.GetLogger().Error("Failed to find opened db connection")
		return errors.New("Failed to find opened db connection")
	}
	return nil
}

func (userRepository *UserRepository) Create(user *models.User) *models.User {
	if userRepository.checkDBConnection() != nil {
		return nil
	}

	response := userRepository.db.Create(&user)
	if response.Error != nil {
		application.GetLogger().Errorf("failed to created user with error %s", response.Error.Error())
		return nil
	}
	return user
}

func (userRepository *UserRepository) List(listRequest *utils.ListRequest) []models.User {
	if userRepository.checkDBConnection() != nil {
		return nil
	}

	var userList []models.User

	sortField := fmt.Sprintf(" \"%s\" %s ", listRequest.OrderBy, listRequest.Order)
	db := utils.GenerateSqlCondition(userRepository.db, listRequest.Filters).
		Order(sortField).
		Offset(listRequest.Skip).
		Limit(listRequest.PageSize).
		Find(&userList)

	if db.Error != nil {
		application.GetLogger().Error("Failed to get users with err ==> ", db.Error.Error())
		return nil
	}

	return userList
}

func (userRepository *UserRepository) Find(user *models.User) *models.User {
	if userRepository.checkDBConnection() != nil {
		return nil
	}

	foundUser := models.NewUser()
	db := userRepository.db.Where(user)

	response := db.First(&foundUser)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to find user db %s", response.Error.Error())
		return nil
	}
	return foundUser
}

func (userRepository *UserRepository) Delete(user *models.User) error {
	if userRepository.checkDBConnection() != nil {
		return nil
	}
	response := userRepository.db.Delete(user)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to delete user %s", response.Error.Error())
		return response.Error
	}
	return nil
}
