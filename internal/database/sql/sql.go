package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"main/internal/config"
)

func Connect(conf *config.Configuration) *sql.DB {

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		conf.Sql.Host, conf.Sql.User, conf.Sql.Password, conf.Sql.Port, conf.Sql.Database)

	// Create connection pool
	db, err := sql.Open(conf.Sql.Name, connString)
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
