package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
	"log"
)

func Connect() *sql.DB {

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		viper.Get("sql.host"), viper.Get("sql.user"), viper.Get("sql.pass"), viper.Get("sql.port"), viper.Get("sql.db"))
	driverName := fmt.Sprintf("%s", viper.Get("sql.name"))

	// Create connection pool
	db, err := sql.Open(driverName, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected")
	return db
}
