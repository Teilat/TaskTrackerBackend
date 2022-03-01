package sql

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
	"os"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/sqlserver"
	"main/internal/database/sql/models"
)

type DatabaseProvider struct {
	Db       *reform.DB
	DbLogger *log.Logger
}

func Init() DatabaseProvider {
	// reform-db -db-driver=sqlserver -db-source=server=localhost;userid=admin;password=Kot_456789;port=1433;database=mainDb;  init -gofmt=false

	// Create logger for sql operations
	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		viper.Get("sql.host"), viper.Get("sql.user"), viper.Get("sql.pass"), viper.Get("sql.port"), viper.Get("sql.db"))
	driverName := fmt.Sprintf("%s", viper.Get("sql.name"))

	// Create connection pool
	sqlDb, err := sql.Open(driverName, connString)
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

	return DatabaseProvider{
		Db:       db,
		DbLogger: logger,
	}
}

func (DbProvider DatabaseProvider) CreateNewTag(tagName, tagColor string) {
	newUuid := uuid.New()
	s := &models.Tags{
		Id:       newUuid,
		TagName:  tagName,
		TagColor: tagColor,
	}

	DbProvider.Db.Save(s)
}
func (DbProvider DatabaseProvider) GetAllTags() {

	from, err := DbProvider.Db.SelectAllFrom(TagsTable, "")
	if err != nil {
		return
	}

}
