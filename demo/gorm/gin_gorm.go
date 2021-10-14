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

// Optional GinConfig
func createGinConfig() ginoperator.Config {
	return ginoperator.Config{
		// Middleware Example:
		Middleware: []gin.HandlerFunc{
			func(c *gin.Context) {
				c.Header("X-MiddleWare", "1")
				c.Next()
			},
		},
	}
}

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

	// Init gin
	g := gin.Default()

	// new Crudist instance (ginConfig is optional)
	server := ginoperator.Gin(g, createGinConfig())
	storage := gormstorage.Gorm(db)
	c := crudist.New(server, storage)

	// crudist Handler for user model
	crudist.Handle(c, "user/", &models.GormUser{})

	// Start gin server
	g.Run("localhost:3000")
}
