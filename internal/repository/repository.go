package repository

import (
	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/models"
)

type DatabaseRepo interface {
	// user
	InsertUser(user *models.User) (string, error)
	GetUser(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	// contact
	InsertContact(contact *models.Contact) (string, error)
	GetContacts(args models.GetAllUsers) ([]models.Contact, error)
	GetContact(id string) (models.Contact, error)
	UpdateContact(contact *models.Contact) error
	DeleteContact(id, userid string) error

	// session
	CreateSession(session *models.Session) (models.Session, error)
	GetSession(id uuid.UUID) (models.Session, error)
}
