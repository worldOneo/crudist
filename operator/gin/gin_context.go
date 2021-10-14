package ginoperator

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/worldOneo/crudist"
)

// Context crudist.Context for Gin
type Context struct {
	ctx *gin.Context
}

// NewGinContext creates a new context for gin
func NewGinContext(ctx *gin.Context) crudist.Context {
	return &Context{ctx}
}

// Param reats a path parameter
func (g *Context) Param(name string) string {
	return g.ctx.Param(name)
}

// JSONBody reads the request body and parses it
func (g *Context) JSONBody(obj interface{}) error {
	return json.NewDecoder(g.ctx.Request.Body).Decode(obj)
}

// JSON sends code and obj as json response
func (g *Context) JSON(code int, obj interface{}) error {
	g.ctx.JSON(code, obj)
	return nil
}
