package fiberoperator

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/worldOneo/crudist"
)

// Operator type for Crudists gin implementation
type Operator struct {
	server     *fiber.App
	Middleware []fiber.Handler
}

// Config type to provide additional data for crudist gin
type Config struct {
	Middleware []fiber.Handler
}

// Fiber creates a new gin operator for crudist
func Fiber(server *fiber.App, configs ...Config) *Operator {
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
	g.server.Add(httpMethod, path, g.toFiberHandler(handler...)...)
}

func (g *Operator) toContext(ctx *fiber.Ctx) crudist.Context {
	return NewFiberContext(ctx)
}

func (g *Operator) toFiberHandler(handlers ...crudist.HandlerFunc) []fiber.Handler {
	new := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		new[i] = func(ctx *fiber.Ctx) error {
			return handler(g.toContext(ctx))
		}
	}
	return append(g.Middleware, new...)
}
