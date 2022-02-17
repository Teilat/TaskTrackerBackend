package main

import (
	"fmt"
	"main/database/sql"
	"main/server"
	_ "net/http/pprof"
)

func main() {
	sql.Connect()
	server.Init()

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
