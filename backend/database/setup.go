package database

import (
	"fmt"
	"spotify-clone/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
    Host     = "localhost"
    User     = "damarpradnyana"
    Password = "secret"
    DBName   = "spotifyclone"
    Port     = 5432
)

func ConnectDB() *gorm.DB{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", 
		Host,
		User,
		Password,
		DBName,
		Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	helpers.CheckPanic(err)


	return db;
}