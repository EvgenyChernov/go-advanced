package db

import (
	"app/adv-http/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDB(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{DB: db}
}
