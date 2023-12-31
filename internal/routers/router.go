package routers

import (
	"backEnd_Coffeshop/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(config.CorsConfig))

	products(router, db)
	categories(router, db)
	users(router, db)
	auth(router, db)

	return router
}
