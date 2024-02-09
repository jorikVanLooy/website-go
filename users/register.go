package users

import (
	"database/sql"
	"log"
)

func Register(user string, password string, email string, phone string, db string) int64 {
	database, err := sql.Open("sqlite3", db)
	if err != nil {
		log.Fatal(err)
	}
	res, err := database.Exec("insert into users (username, password, email, phone) values(?,?,?,?);", user, password, email, phone)
	if err != nil {
		log.Fatal(err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		log.Fatal(err)
	} else {
		return id
	}
	return 0
}
