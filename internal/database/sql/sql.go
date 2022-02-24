package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/sqlserver"
	"log"
	"main/internal/config"
	"os"
)

type Tags struct {
}

func Init(conf *config.Configuration) (*reform.DB, *log.Logger) {

	// Create logger for sql operations
	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		conf.Sql.Host, conf.Sql.User, conf.Sql.Password, conf.Sql.Port, conf.Sql.Database)

	// Create connection pool
	sqlDb, err := sql.Open(conf.Sql.Name, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = sqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected")

	// Create Db using reform ORM
	db := reform.NewDB(sqlDb, sqlserver.Dialect, reform.NewPrintfLogger(logger.Printf))

	return db, logger
}

func GetAllTags(db *reform.DB) {

}
