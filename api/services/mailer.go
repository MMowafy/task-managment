package services

import (
	"task-managment/api/models"
)

type MailerService struct {
}

func NewMailerService() *MailerService {
	return &MailerService{
	}
}

func (MailerService *MailerService) SendEmail(task *models.Task) {
	// TODO: Prepare data and publish it to message queue to be consumed by mailer cli (another microservice)
}
