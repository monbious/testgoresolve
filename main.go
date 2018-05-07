package main

import (
	"testgoresolve/server"
	"testgoresolve/dbUtils"
)

func main() {
	dbUtils.CheckTables()
	server.RunServer()
}