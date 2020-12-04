package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

type Person struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age uint8 `json:"age"`
	Username string `json:"username" binding:"required"`
}

func main(){
	var p = Person{Name:"Ram",Email:"ram@ayodhya.in",Username:"ram"}
	fmt.Println(p.Name)
	var t = reflect.TypeOf(Person{})
	fieldName, _ := t.FieldByName("Name")
	fmt.Println(fieldName)
	fmt.Println(fieldName.Tag)
	v,ok := fieldName.Tag.Lookup("json")
	fmt.Println(v,ok)
	v,ok = fieldName.Tag.Lookup("binding")
	fmt.Println(v,ok)

	b, err := json.Marshal(p)
	fmt.Printf("%s",b)
	fmt.Println(err)

}