package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/worldOneo/crudist"
	"github.com/worldOneo/crudist/demo/models"
	fiberoperator "github.com/worldOneo/crudist/operator/fiber"
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
	db.Migrator().AutoMigrate(&models.GormUser{})

	// Init Fiber
	g := fiber.New()
	g.Use(fiberlogger.New())

	// Setup operator and storage crudist
	server := fiberoperator.Fiber(g)
	storage := gormstorage.Gorm(db)
	c := crudist.New(server, storage)

	// crudist Handler for user model
	crudist.Handle(c, "user/", &models.GormUser{})

	// Start fiber server
	g.Listen("localhost:3000")
}
