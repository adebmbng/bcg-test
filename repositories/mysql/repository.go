package mysql

import (
	"github.com/adebmbng/bcg-test/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetInventoriesBySKUs(q []string) ([]*entities.Inventory, error)

	GetPromoBySKUs(q []string) ([]*entities.Promotion, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
