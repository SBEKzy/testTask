package repository

import (
	"fmt"
	"github.com/SBEKzy/testTask/config"
	"github.com/SBEKzy/testTask/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type Repository interface {
	SaveRequest(request, response string) error
}

type repo struct {
	DB *gorm.DB
}

func New(conf *config.Config) Repository {
	db := connectDB(conf)
	return &repo{
		DB: db,
	}
}

func connectDB(conf *config.Config) *gorm.DB {

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&model.Save{},
	)

	return db
}
