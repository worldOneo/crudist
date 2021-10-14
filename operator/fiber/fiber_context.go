package fiberoperator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/worldOneo/crudist"
)

// Context crudist.Context for Gin
type Context struct {
	ctx *fiber.Ctx
}

// NewFiberContext creates a new context for gin
func NewFiberContext(ctx *fiber.Ctx) crudist.Context {
	return &Context{ctx}
}

// Param reats a path parameter
func (g *Context) Param(name string) string {
	return g.ctx.Params(name)
}

// JSONBody reads the request body and parses it
func (g *Context) JSONBody(obj interface{}) error {
	return g.ctx.BodyParser(obj)
}

// JSON sends code and obj as json response
func (g *Context) JSON(code int, obj interface{}) error {
	return g.ctx.Status(code).JSON(obj)
}
