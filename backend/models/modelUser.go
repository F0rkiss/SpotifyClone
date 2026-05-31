package models

import (
	"time"
)

type User struct {
	Id             string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email          string    `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	Username       string    `gorm:"type:varchar(255);unique_index;not null" json:"username"`
	HashedPassword string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// ga perlu di isi pas register
	Age       string     `gorm: "type:varchar(50)" json:"age"`
	Gender    string     `gorm: "type:varchar(50)" json:"gender"`
	BirthDate *time.Time `gorm: "type:date" json:"birth_date`
	Country   string     `gorm:"type:varchar(100)" json:"country"`
}
