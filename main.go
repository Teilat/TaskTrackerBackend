package main

import (
	"fmt"
	"main/config"
	"main/database/sql"
	"main/server"
	_ "main/server/api/v1/models"
	_ "net/http/pprof"
)

// @host      localhost:8080
// @BasePath  /api/v1
// @Title     Application Api
// @Version   1.0
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