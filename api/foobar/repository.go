package foobar

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// Repository interface
type Repository interface {
	Fetch() (*gorm.DB, *[]models.Foobar, error)
	GetByID(id uuid.UUID) (*models.Foobar, error)
	Update(vT *models.Foobar) error
	Store(vT *models.Foobar) error
	Delete(id uuid.UUID) error
}
