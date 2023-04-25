package services

import "github.com/farakh-shahid/CRUD-APP/models"

type UserServiceInterface interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(string, *models.User) error
	DeleteUser(*string) error
}
