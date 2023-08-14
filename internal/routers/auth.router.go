package routers

import (
	"backEnd_Coffeshop/internal/handlers"
	"backEnd_Coffeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	repo := repositories.NewUsers(d)
	handler := handlers.NewAuth(repo)

	route.POST("/login", handler.Login)

}
