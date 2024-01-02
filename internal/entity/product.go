package entity

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type (
	Product struct {
		ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid()"`
		Name        string
		Description string
		Price       float32
		CreatedAt   int
		UpdatedAt   int
		DeletedAt   soft_delete.DeletedAt `gorm:"softDelete:nano"`
	}
)
