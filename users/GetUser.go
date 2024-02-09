package users

import (
	"database/sql"
	"fmt"
	"log"
)

func GetUser(user string, db string) string {
	database, err := sql.Open("sqlite3", db)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Query(fmt.Sprintf("select username from users where id = '%s'", user))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		rows.Scan(&username)
		return username
	}
	return "anonymous"
}
