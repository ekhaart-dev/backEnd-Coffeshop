package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Pgdb() (*sqlx.DB, error) {
	// host := viper.GetString("database.host")
	// port := viper.GetString("database.port")
	// user := viper.GetString("database.user")
	// password := viper.GetString("database.pass")
	// dbName := viper.GetString("database.name")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// host := "127.0.0.1"
	// port := 5433
	// user := "ekha"
	// password := "@0okmnji9!"
	// dbName := "coffeshop_database"

	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbName)

	return sqlx.Connect("postgres", config)

}
