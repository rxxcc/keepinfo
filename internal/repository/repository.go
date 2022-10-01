package repository

import "github.com/inuoshios/keepinfo/internal/models"

type Repository interface {
	InsertUser(user models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
