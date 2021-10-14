package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/worldOneo/crudist"
	"github.com/worldOneo/crudist/demo/models"
	ginoperator "github.com/worldOneo/crudist/operator/gin"
	gormstorage "github.com/worldOneo/crudist/storage/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// SQL Connection
	sql := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", "root", "1234", "localhost", "test"))

	// Init gorm
	db, err := gorm.Open(sql, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// Create tables
	db.Migrator().AutoMigrate(&models.GormPost{}, &models.GormUser{})

	// Init gin
	g := gin.Default()

	// new Crudist instance (ginConfig is optional)
	server := ginoperator.Gin(g)
	userStorage := gormstorage.Gorm(db, gormstorage.Config{
		Inceptors: gormstorage.Inceptors{
			GetByID: []gormstorage.Inceptor{
				func(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
					return tx.Preload("GormPosts"), false
				},
			},
		},
	})

	userJoin := func(model interface{}, tx *gorm.DB) (*gorm.DB, bool) {
		return tx.Joins("GormUser"), false
	}

	postStorage := gormstorage.Gorm(db, gormstorage.Config{
		Inceptors: gormstorage.Inceptors{
			Get: []gormstorage.Inceptor{userJoin},
			GetByID: []gormstorage.Inceptor{userJoin},
		},
	})

	users := crudist.New(server, userStorage)
	posts := crudist.New(server, postStorage)


	// crudist Handler for user model
	crudist.Handle(users, "user/", &models.GormUser{})
	crudist.Handle(posts, "post/", &models.GormPost{})

	// Start gin server
	g.Run("localhost:3000")
}
