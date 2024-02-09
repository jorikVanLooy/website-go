package main

import (
	routers "website/routers"

	_ "github.com/mattn/go-sqlite3"
)

const db string = "users.db"

func main() {

	router := routers.Routers(db)

	router.Run()
}
