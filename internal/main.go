package main

import (
	"fmt"
	"main/internal/config"
	"main/internal/database/sql"
	"main/internal/server"
	_ "main/internal/server/api/v1/models"
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
		fmt.Printf("Database error:%s", err)
		return
	}

	// setting up server
	server.Init()
}
