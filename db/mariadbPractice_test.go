package db
import{
	"testing"
}
func Test_Open_mariaDB() *sql.DB {
	db, err := sql.Open("mysql", "root:qeek1688@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func Test_Close_mariaDB(db *sql.DB) {
	defer db.Close()
}
func Test_DB_connect(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		println("Fail to connect")
	}
}
func Test_Fetch_data_by_name(db *sql.DB, name string) {
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
