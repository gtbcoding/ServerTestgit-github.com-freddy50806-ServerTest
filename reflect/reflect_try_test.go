package StructBind

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	UserID    int    `json:"UserID"`
	Username  string `json:"Username"`
	Useremail string `json:"Usermail"`
}
type Game1 struct {
	Game1ID     int    `json:"Game1ID"`
	Game1Level  int    `json:"Game1LV"`
	Game1avatar string `json:"Game1character"`
	Game1STR    int    `json:"Game1STR"`
	Game1AGI    int    `json:"Game1AGI"`
	Game1INT    int    `json:"Game1INT"`
}
type Game2 struct {
	Game2ID     int    `json:"Game2ID"`
	Game2Level  int    `json:"Game2LV"`
	Game2avatar string `json:"Game2character"`
	Game2Item1  string `json:"Game2Item1"`
	Game2Item2  string `json:"Game2Item2"`
	Game2Item3  string `json:"Game2Item3"`
}
type outputdata struct {
	UserID      int    `json:"UserID"`
	Username    string `json:"Username"`
	Game1ID     int    `json:"Game1ID"`
	Game1Level  int    `json:"Game1LV"`
	Game1avatar string `json:"Game1character"`
}

func Test_StructBind(t *testing.T) {
	gary := User{777, "Gary", "gogo@qnap.com"}
	//commander := Game1{2131, 9, "Avatar_gary", 6, 7, 4}
	//gary_data := outputdata{}
	//err := StructBind(gary, commander, gary_data)
	fmt.Printf("User StructTag=%T\n", gary)
	println("User typeOf=", reflect.TypeOf(gary))
	/*
		if err == true {
			t.Log("Binding success")
		} else if err == false {
			t.Error("Binding failed")
		}
	*/
	// GXA := "123"
	// var i int64 = 1234
	// var j int = 23254
	// if reflect.TypeOf(&gary).Elem().Kind().String() == "struct" {
	// 	println("true")
	// }
	// println(reflect.TypeOf(&i).Elem().Kind().String())
	// println(reflect.TypeOf(&j).Elem().Kind().String())
	// println(reflect.TypeOf(&GXA).Elem().Kind().String())
	// println(reflect.TypeOf(&gary))
	// println(reflect.TypeOf(&gary).Elem().Kind())
	// fmt.Printf("GG=%s\n", reflect.TypeOf(gary).Kind().String())
	// println("G2", reflect.TypeOf(gary).Kind().String())
	s := reflect.ValueOf(gary)
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
