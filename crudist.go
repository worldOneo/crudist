package crudist

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Context is custom for crudist handlers
type Context interface {
	JSONBody(obj interface{}) error
	JSON(code int, obj interface{}) error
	Param(name string) string
}

// HandlerFunc for Crudist
type HandlerFunc func(ctx Context) error

// Crudist base interface definition
type Crudist interface {
	// DB returns the db for the crudist instance
	DB() *gorm.DB
	// Get registers new handlers for http get method
	Get(path string, handler ...HandlerFunc)
	// Get registers new handlers for http Post method
	Post(path string, handler ...HandlerFunc)
	// Get registers new handlers for http Patch method
	Patch(path string, handler ...HandlerFunc)
	// Get registers new handlers for http Delete method
	Delete(path string, handler ...HandlerFunc)
}

// Gin creates a new gin operator for crudist
func Gin(server *gin.Engine, db *gorm.DB) *GinOperator {
	return &GinOperator{server, db}
}

// ErrorBadRequest when JSON parsing failed
var ErrorBadRequest error = fmt.Errorf("Bad Request")
// ErrorInternalServer when the DB couldn't 
var ErrorInternalServer error = fmt.Errorf("Internal Server Error")

// Handle creates handler for GET POST PATCH and DELETE operation for your model
func Handle(crudist Crudist, path string, model interface{}) {
	modelType := reflect.TypeOf(model)
	sliceType := reflect.SliceOf(modelType)
	newModel := func() interface{} {
		return reflect.New(modelType).Interface()
	}

	newModels := func() interface{} {
		return reflect.New(sliceType).Interface()
	}

	crudist.Get(path, func(ctx Context) error {
		models := newModels()
		fmt.Print(reflect.TypeOf(models))
		err := crudist.DB().Find(models).Error
		if err != nil {
			return err
		}
		return ctx.JSON(200, models)
	})

	crudist.Get(path+":id/", func(ctx Context) error {
		stringID := ctx.Param("id")
		id, err := strconv.Atoi(stringID)
		if err != nil {
			return ErrorBadRequest
		}
		model := newModel()
		err = crudist.DB().Find(model, id).Error
		if err != nil {
			return err
		}
		return ctx.JSON(200, model)
	})

	crudist.Post(path, func(ctx Context) error {
		model := newModel()
		err := ctx.JSONBody(model)
		if err != nil {
			log.Print(err)
			return ErrorBadRequest
		}
		err = crudist.DB().Create(model).Error
		if err != nil {
			return ErrorInternalServer
		}
		return ctx.JSON(200, model)
	})

	crudist.Delete(path, func(ctx Context) error {
		model := newModel()
		err := ctx.JSONBody(model)
		if err != nil {
			log.Print(err)
			return ErrorBadRequest
		}
		err = crudist.DB().Delete(model).Error
		if err != nil {
			return ErrorInternalServer
		}
		return ctx.JSON(200, model)
	})

	crudist.Patch(path, func(ctx Context) error {
		model := newModel()
		err := ctx.JSONBody(model)
		if err != nil {
			return ErrorBadRequest
		}
		err = crudist.DB().Save(model).Error
		if err != nil {
			return ErrorInternalServer
		}
		return ctx.JSON(200, model)
	})
}
