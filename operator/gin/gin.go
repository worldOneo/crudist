package ginoperator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldOneo/crudist"
)

// Operator type for Crudists gin implementation
type Operator struct {
	server     *gin.Engine
	Middleware []gin.HandlerFunc
}

// Config type to provide additional data for crudist gin
type Config struct {
	Middleware gin.HandlersChain
}

// Gin creates a new gin operator for crudist
func Gin(server *gin.Engine, configs ...Config) *Operator {
	config := Config{}
	if len(configs) > 0 {
		config = configs[0]
	}
	return &Operator{server, config.Middleware}
}

// Get registers new handlers for http get method
func (g *Operator) Get(path string, handler ...crudist.HandlerFunc) {
	g.handle(http.MethodGet, path, handler)
}

// Post registers new handlers for http post method
func (g *Operator) Post(path string, handler ...crudist.HandlerFunc) {
	g.handle(http.MethodPost, path, handler)
}

// Patch registers new handlers for http patch method
func (g *Operator) Patch(path string, handler ...crudist.HandlerFunc) {
	g.handle(http.MethodPatch, path, handler)
}

// Delete registers new handlers for http delete method
func (g *Operator) Delete(path string, handler ...crudist.HandlerFunc) {
	g.handle(http.MethodDelete, path, handler)
}

func (g *Operator) handle(httpMethod, path string, handler []crudist.HandlerFunc) {
	g.server.Handle(httpMethod, path, g.toGinHandler(handler...)...)
}

func (g *Operator) toContext(ctx *gin.Context) crudist.Context {
	return NewGinContext(ctx)
}

func (g *Operator) toGinHandler(handlers ...crudist.HandlerFunc) []gin.HandlerFunc {
	new := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		new[i] = func(ctx *gin.Context) {
			err := handler(g.toContext(ctx))
			if err != nil {
				ctx.Error(err)
			}
		}
	}
	return append(g.Middleware, new...)
}
