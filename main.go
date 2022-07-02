package main

import (
	"fmt"
	"main/config"
	"main/db/sql"
	"main/server"
	_ "main/server/api/v1/models"
	_ "net/http/pprof"
)

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
