package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    Username  string         `gorm:"unique;not null" json:"username" validate:"required,min=3,max=50"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}