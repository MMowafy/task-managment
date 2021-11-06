package infrastructure

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"task-managment/api/application"
	"task-managment/api/models"
	"task-managment/api/utils"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() *NotificationRepository {
	db, _ := application.GetPostgresConnectionByName("appdb")
	NotificationRepository := &NotificationRepository{
		db,
	}
	NotificationRepository.db.LogMode(application.GetConfig().GetBool("app.db_logs"))
	return NotificationRepository
}

func (notificationRepository *NotificationRepository) checkDBConnection() error {
	if notificationRepository.db == nil {
		application.GetLogger().Error("Failed to find opened db connection")
		return errors.New("Failed to find opened db connection")
	}
	return nil
}

func (notificationRepository *NotificationRepository) Create(Notification *models.Notification) *models.Notification {
	if notificationRepository.checkDBConnection() != nil {
		return nil
	}

	response := notificationRepository.db.Create(&Notification)
	if response.Error != nil {
		application.GetLogger().Errorf("failed to created Notification with error %s", response.Error.Error())
		return nil
	}
	return Notification
}

func (notificationRepository *NotificationRepository) Find(Notification *models.Notification) *models.Notification {
	if notificationRepository.checkDBConnection() != nil {
		return nil
	}

	foundNotification := models.NewNotification()
	db := notificationRepository.db.Where(Notification)

	response := db.First(&foundNotification)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to find Notification db %s", response.Error.Error())
		return nil
	}
	return foundNotification
}

func (notificationRepository *NotificationRepository) List(listRequest *utils.ListRequest) []models.Notification {
	if notificationRepository.checkDBConnection() != nil {
		return nil
	}

	var NotificationList []models.Notification

	sortField := fmt.Sprintf(" \"%s\" %s ", listRequest.OrderBy, listRequest.Order)
	db := utils.GenerateSqlCondition(notificationRepository.db, listRequest.Filters).
		Order(sortField).
		Offset(listRequest.Skip).
		Limit(listRequest.PageSize).
		Find(&NotificationList)

	if db.Error != nil {
		application.GetLogger().Error("Failed to get Notification with err ==> ", db.Error.Error())
		return nil
	}

	return NotificationList
}
