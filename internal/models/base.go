package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;type:uuid;default:public.uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedBy string         `json:"created_by"`
	UpdatedBy string         `json:"updated_by"`
}
