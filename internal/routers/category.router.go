package routers

import (
	"backEnd_Coffeshop/internal/handlers"
	"backEnd_Coffeshop/internal/middleware"
	"backEnd_Coffeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func categories(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/categories")

	// dependency injection
	repo := repositories.NewCategory(d)
	handler := handlers.NewCategory(repo)

	route.GET("/", handler.GetCategories)
	route.POST("/", middleware.AuthJwt("admin"), handler.CreateCategory)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.DeleteCategory)
}
