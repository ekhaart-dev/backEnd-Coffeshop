package routers

import (
	"backEnd_Coffeshop/internal/handlers"
	"backEnd_Coffeshop/internal/middleware"
	"backEnd_Coffeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/products")

	// dependency injection
	repo := repositories.NewProduct(d)
	handler := handlers.NewProductHandler(repo)

	route.GET("/", handler.GetProducts)
	route.POST("/", middleware.AuthJwt("admin"), middleware.UploadFile, handler.CreateProduct)
	route.PUT("/:id", middleware.AuthJwt("admin"), middleware.UploadFile, handler.UpdateProduct)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.DeleteProduct)
}
