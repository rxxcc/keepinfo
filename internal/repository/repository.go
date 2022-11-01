package repository

import "github.com/inuoshios/keepinfo/internal/models"

type DatabaseRepo interface {
	InsertUser(user *models.User) (string, error)
	GetUserbyEmail(email string) (*models.User, error)
}
