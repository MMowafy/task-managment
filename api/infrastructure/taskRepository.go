package infrastructure

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"task-managment/api/application"
	"task-managment/api/models"
	"task-managment/api/utils"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() *TaskRepository {
	db, _ := application.GetPostgresConnectionByName("appdb")
	TaskRepository := &TaskRepository{
		db,
	}
	TaskRepository.db.LogMode(application.GetConfig().GetBool("app.db_logs"))
	return TaskRepository
}

func (taskRepository *TaskRepository) checkDBConnection() error {
	if taskRepository.db == nil {
		application.GetLogger().Error("Failed to find opened db connection")
		return errors.New("Failed to find opened db connection")
	}
	return nil
}

func (taskRepository *TaskRepository) Create(task *models.Task) *models.Task {
	if taskRepository.checkDBConnection() != nil {
		return nil
	}

	response := taskRepository.db.Create(&task)
	if response.Error != nil {
		application.GetLogger().Errorf("failed to created Task with error %s", response.Error.Error())
		return nil
	}
	return task
}

func (taskRepository *TaskRepository) Find(task *models.Task) *models.Task {
	if taskRepository.checkDBConnection() != nil {
		return nil
	}

	foundTask := models.NewTask()
	db := taskRepository.db.Where(task)

	response := db.First(&foundTask)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to find task db %s", response.Error.Error())
		return nil
	}
	return foundTask
}

func (taskRepository *TaskRepository) List(listRequest *utils.ListRequest) []models.Task {
	if taskRepository.checkDBConnection() != nil {
		return nil
	}

	var taskList []models.Task

	sortField := fmt.Sprintf(" \"%s\" %s ", listRequest.OrderBy, listRequest.Order)
	db := utils.GenerateSqlCondition(taskRepository.db, listRequest.Filters).
		Order(sortField).
		Offset(listRequest.Skip).
		Limit(listRequest.PageSize).
		Find(&taskList)

	if db.Error != nil {
		application.GetLogger().Error("Failed to get task with err ==> ", db.Error.Error())
		return nil
	}

	return taskList
}

func (taskRepository *TaskRepository) Delete(task *models.Task) error {
	if taskRepository.checkDBConnection() != nil {
		return nil
	}
	response := taskRepository.db.Delete(task)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to delete task %s", response.Error.Error())
		return response.Error
	}
	return nil
}

func (taskRepository *TaskRepository) FindOverlappingTasks(task *models.Task) *models.Task {
	if taskRepository.checkDBConnection() != nil {
		return nil
	}

	foundTask := models.NewTask()
	db := taskRepository.db.Where("user_id = ?", task.UserId).Where("(start_at <= ? AND end_at > ?) OR (start_at <= ? AND end_at >= ?) OR (start_at >= ? AND end_at <= ?)", task.StartAt, task.StartAt, task.EndAt, task.EndAt, task.StartAt, task.EndAt)

	response := db.First(&foundTask)
	if response.Error != nil {
		application.GetLogger().Errorf("Failed to find task db %s", response.Error.Error())
		return nil
	}
	return foundTask
}
