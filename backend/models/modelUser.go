package models

import (
	"time"
)

type User struct { 
	Id string `gorm: "type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email string `gorm:"type:varchar(255);unique_index" json:"email"`
	HashedPassword string    `gorm:"type:varchar(255)" json:"-"`
    CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	Username string `gorm: "type:varchar(255)`
	Age string `gorm: "type:varchar(255)`
	Gender string `gorm: "type:varchar(255)`
	Status string `gorm: "type:varchar(255)`
}