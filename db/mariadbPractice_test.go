package db

import "testing"
import "fmt"

func Test_mariaDB(t *testing.T) {

	//test 1
	db, err := Open_mariaDB()
	if err != nil {
		t.Error("fail to open db")
	} else {
		t.Log("Open_mariaDB success")
	}
	//test 2
	err = DB_connect(db)
	if err != nil {
		t.Error("DB_connect failed")
	} else {
		t.Log("DB_connect success")
	}
	//test 3
	user1, err := Fetch_data_by_name(db, "freddy")
	fmt.Println(user1)

	if user1 == nil {
		t.Error("Fetch_data_by_name fialed")
	} else {
		t.Log("Fetch_data_by_name success")
	}

}
