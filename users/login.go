package users

import (
	"database/sql"
	"fmt"
	"log"
)

func Login(user string, inputPassword string, db string) (bool, string) {
	database, err := sql.Open("sqlite3", db)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Query(fmt.Sprintf("select id, password from users where username = '%s'", user))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int64
		var password string
		err := rows.Scan(&id, &password)
		if err != nil {
			log.Fatal(err)
		}
		return password == inputPassword, fmt.Sprintf("%d", id)
	}
	return false, "0"
}
