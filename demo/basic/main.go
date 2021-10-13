package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/worldOneo/crudist"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	BaseModel
	Username string `json:"username" gorm:"size:100"`
	Password string `json:"password" gorm:"size:128"`
}

func createGinConfig() crudist.GinConfig {
	return crudist.GinConfig{
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

	// Create databases
	db.Migrator().AutoMigrate(&User{})

	// Init gin
	g := gin.Default()

	// new Crudist instance (ginConfig is optional)
	c := crudist.Gin(g, db, createGinConfig())

	// crudist Handler for user model
	crudist.Handle(c, "user/", &User{})

	// Start gin server
	g.Run("localhost:3000")
}
