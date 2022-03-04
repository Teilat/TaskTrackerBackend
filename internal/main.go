package main

import (
	"main/internal/api"
	_ "main/internal/api/v1/models"
	"main/internal/config"
	"main/internal/database/sql"
	_ "net/http/pprof"
)

// @host      localhost:8080
// @BasePath  /api/v1
func main() {

	// getting config from config file
	config.GetConf()

	// setting up db connection
	_, err := sql.Init()
	if err != nil {
		return
	}

	// setting up api
	api.Init()
}
