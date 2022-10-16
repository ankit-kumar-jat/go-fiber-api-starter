package db

import (
	"fmt"
	"log"

	"example.com/go-api-starter/internal/config"
	"example.com/go-api-starter/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

func Init(c *config.Config) {
    db := DB

    url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    // advance configuration visit https://github.com/antonioalfa22/go-rest-template/blob/master/internal/pkg/db/database.go
    // db.LogMode(false)
	// db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	// db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	// db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
    
    DB = db
    migration()

}

func migration() {
	DB.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}