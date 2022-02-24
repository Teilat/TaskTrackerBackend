package main

import (
	"fmt"
	"log"
	"main/internal/api"
	"main/internal/config"
	"main/internal/database/sql"
	_ "net/http/pprof"
)

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	var conf config.Configuration
	conf.GetConf()

	db := sql.Connect(&conf)
	sql.GetAllTags(db)
	api.Init()

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
