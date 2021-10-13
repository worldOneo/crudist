package crudist

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GinOperator type for Crudists gin implementation
type GinOperator struct {
	server *gin.Engine
	db *gorm.DB
}


func (g *GinOperator) DB() *gorm.DB {
	return g.db
}

func (g *GinOperator) Get(path string, handler ...HandlerFunc) {
	g.server.GET(path, g.toGinHandler(handler...)...)
}

func (g *GinOperator) Post(path string, handler ...HandlerFunc) {
	g.server.POST(path, g.toGinHandler(handler...)...)
}

func (g *GinOperator) Patch(path string, handler ...HandlerFunc) {
	g.server.PATCH(path, g.toGinHandler(handler...)...)
}

func (g *GinOperator) Delete(path string, handler ...HandlerFunc) {
	g.server.DELETE(path, g.toGinHandler(handler...)...)
}



func (g *GinOperator) toGinHandler(handlers ...HandlerFunc) []gin.HandlerFunc {
	new := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		new[i] = func(ctx *gin.Context) {
			err := handler(g.toContext(ctx))
			if err != nil {
				ctx.Error(err)
			}
		}
	}
	return new
}