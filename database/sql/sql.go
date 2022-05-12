package sql

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/sqlserver"
)

var instance *DatabaseProvider
var once sync.Once

type DatabaseProvider struct {
	DB       *reform.DB
	DbLogger *log.Logger
}

func Init() (*DatabaseProvider, error) {
	// cd database/sql/models
	// reform-db -db-driver=sqlserver -db-source=server=localhost;userid=admin;password=Kot_456789;port=1433;database=mainDb; init -gofmt=false

	// Create logger for sql operations
	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		viper.Get("sql.host"), viper.Get("sql.user"), viper.Get("sql.pass"), viper.Get("sql.port"), viper.Get("sql.db"))
	driverName := fmt.Sprintf("%s", viper.Get("sql.name"))

	// Create connection pool
	sqlDb, err := sql.Open(driverName, connString)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = sqlDb.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	logger.Println("Connected")

	// Create Db using reform ORM
	db := reform.NewDB(sqlDb, sqlserver.Dialect, reform.NewPrintfLogger(logger.Printf))

	instance = &DatabaseProvider{
		DB:       db,
		DbLogger: logger,
	}

	return instance, nil
}

func GetDb() *DatabaseProvider {
	once.Do(func() {
		var err error
		instance, err = Init()
		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}