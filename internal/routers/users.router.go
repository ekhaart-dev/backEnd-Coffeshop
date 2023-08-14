package routers

import (
	"backEnd_Coffeshop/internal/handlers"
	"backEnd_Coffeshop/internal/middleware"
	"backEnd_Coffeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUsers(d)
	handler := handlers.NewUsers(repo)

	route.POST("/", handler.PostData)
	route.GET("/", handler.FetchAll)
	route.PUT("/:id", middleware.AuthJwt("admin", "user"), handler.UpdateUsers)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.DeleteUser)
}
