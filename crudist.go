package crudist

import (
	"fmt"
	"reflect"
	"strconv"
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
	// Storage for the storage layer
	Storage() Storage
	// Server for the web framework connection
	Server() Server
	// Config for additional functionality
	Config() Config
}

type baseCrudist struct {
	server  Server
	storage Storage
	config  Config
}

// Storage for the storage layer
func (b *baseCrudist) Storage() Storage {
	return b.storage
}

// Server for the web framework connection
func (b *baseCrudist) Server() Server {
	return b.server
}

// Config for additional functionality
func (b *baseCrudist) Config() Config {
	return b.config
}

// New creates a new Crudist instance from a server and a storage
func New(server Server, storage Storage, config ...Config) Crudist {
	conf := Config{
		IDGenerator: func(id string) (interface{}, error) {
			return strconv.Atoi(id)
		},
		Operations: OperationAll,
	}
	if len(config) > 0 {
		additional := config[0]
		if additional.IDGenerator != nil {
			conf.IDGenerator = additional.IDGenerator
		}
		if additional.Operations != 0 {
			conf.Operations = additional.Operations
		}
	}
	return &baseCrudist{server, storage, conf}
}

// Server interface definition for connecting to the web framework
type Server interface {
	// Get registers new handlers for http get method
	Get(path string, handler ...HandlerFunc)
	// Post registers new handlers for http Post method
	Post(path string, handler ...HandlerFunc)
	// Patch registers new handlers for http Patch method
	Patch(path string, handler ...HandlerFunc)
	// Delete registers new handlers for http Delete method
	Delete(path string, handler ...HandlerFunc)
}

// Storage interface for the storage provider
type Storage interface {
	// Get gets a model slice to populate
	Get(models interface{}) error
	// GetByID gets a model with the specific id
	GetByID(model interface{}, id interface{}) error
	// Create adds a new model
	Create(model interface{}) error
	// Update modifies a model
	Update(model interface{}) error
	// Delete removes a model
	Delete(model interface{}) error
	// DeleteByID deletes a model given an id
	DeleteByID(model, id interface{}) error
}

// CrudOperation type defines one operation of CRUD
type CrudOperation uint64

const (
	// OperationGet allows the get operation to get all
	OperationGet CrudOperation = 1 << iota
	// OperationGetByID allows the get operation to get one item by its ID
	OperationGetByID
	// OperationUpdate allows the Update operation to update an item
	OperationUpdate
	// OperationCreate allows the Create operation to save a new item
	OperationCreate
	// OperationDelete allows the Delete operation to delete an item
	OperationDelete
	// OperationDeleteByID allows the Delete operation to delete an item by its ID
	OperationDeleteByID
	// OperationAll allows all operations
	OperationAll = OperationGet |
		OperationGetByID |
		OperationUpdate |
		OperationCreate |
		OperationDelete |
		OperationDeleteByID
)

// Config for additional data for crudist
type Config struct {
	// IDGenerator to use for the Storage as id
	// id from the path parameter string in the route GET path/:id
	//
	// Default: strconv.Itoa
	IDGenerator func(id string) (interface{}, error)

	// Operations describes the allowed opperations in the
	// the CRUD API with this routes can be added or removed
	// from the API.
	// The value 0 will result in OperationAll as it would
	// be a noop otherwise
	//
	// Default: OperationAll
	Operations CrudOperation
}

// ErrorBadRequest when JSON parsing failed
var ErrorBadRequest error = fmt.Errorf("Bad Request")

// ErrorInternalServer when the storage layer couldn't be accessed correctly
var ErrorInternalServer error = fmt.Errorf("Internal Server Error")

// JSONDoc can be used to create an json document inline
// with a clear purpose
type JSONDoc map[string]interface{}

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

	server := crudist.Server()
	storage := crudist.Storage()
	conf := crudist.Config()

	getBody := func(ctx Context) (interface{}, error) {
		model := newModel()
		err := ctx.JSONBody(model)
		if err != nil {
			return nil, ErrorBadRequest
		}
		return model, nil
	}

	getID := func(ctx Context) (interface{}, error) {
		stringID := ctx.Param("id")
		id, err := conf.IDGenerator(stringID)
		if err != nil {
			return nil, ErrorBadRequest
		}
		return id, nil
	}

	httpError := func(ctx Context, err error) error {
		if err != nil {
			if err == ErrorBadRequest {
				return ctx.JSON(400, JSONDoc{
					"message": "Bad Request",
					"code":    400,
				})
			}
			sendErr := ctx.JSON(500, JSONDoc{
				"message": "Internal Server Error",
				"code":    500,
			})
			if sendErr != nil {
				return sendErr
			}
			return err
		}
		return nil
	}

	createModelRoute := func(operation CrudOperation,
		register func(string, ...HandlerFunc),
		storageFunc func(interface{}) error) {
		if conf.Operations&operation == operation {
			register(path, func(ctx Context) error {
				body, err := getBody(ctx)
				if err != nil {
					return httpError(ctx, err)
				}
				err = storageFunc(body)
				if err != nil {
					return httpError(ctx, err)
				}
				return ctx.JSON(200, body)
			})
		}
	}

	createIDRoute := func(operation CrudOperation,
		register func(string, ...HandlerFunc),
		storageFunc func(interface{}, interface{}) error) {
		if conf.Operations&operation == operation {
			register(path+":id/", func(ctx Context) error {
				id, err := getID(ctx)
				if err != nil {
					return httpError(ctx, err)
				}
				model := newModel()
				err = storageFunc(model, id)
				if err != nil {
					return httpError(ctx, err)
				}
				return ctx.JSON(200, model)
			})
		}
	}

	if conf.Operations&OperationGet == OperationGet {
		server.Get(path, func(ctx Context) error {
			models := newModels()
			err := storage.Get(models)
			if err != nil {
				return err
			}
			return ctx.JSON(200, models)
		})
	}

	createIDRoute(OperationGetByID, server.Get, storage.GetByID)
	createModelRoute(OperationCreate, server.Post, storage.Create)
	createModelRoute(OperationUpdate, server.Patch, storage.Update)
	createModelRoute(OperationDelete, server.Delete, storage.Delete)
	createIDRoute(OperationDeleteByID, server.Delete, storage.DeleteByID)
}
