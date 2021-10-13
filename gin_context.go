package crudist

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// GinContext crudist.Context for Gin
type GinContext struct {
	ctx *gin.Context
}

// NewGinContext creates a new context for gin
func NewGinContext(ctx *gin.Context) Context {
	return &GinContext{ctx}
}

// Param reats a path parameter
func (g *GinContext) Param(name string) string {
	return g.ctx.Param(name)
}

// JSONBody reads the request body and parses it
func (g *GinContext) JSONBody(obj interface{}) error {
	return json.NewDecoder(g.ctx.Request.Body).Decode(obj)
}

// JSON sends code and obj as json response
func (g *GinContext) JSON(code int, obj interface{}) error {
	g.ctx.JSON(code, obj)
	return nil
}
