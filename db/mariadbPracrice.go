package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	email    string
	age      int
}

func Open_mariaDB() *sql.DB {
	db, err := sql.Open("mysql", "root:qeek1688@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func Close_mariaDB(db *sql.DB) {
	defer db.Close()
}
func DB_connect(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		println("Fail to connect")
	}
}
func Fetch_data_by_name(db *sql.DB, name string) {
	var u User
	rows, err := db.Query("SELECT * FROM users WHERE Username=?", name)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(u.Username, u.email, u.age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
