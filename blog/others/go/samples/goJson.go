package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	author string
    country string
    imageLink string
    language string
    link string
    pages uint8
    title string
    year uint16
}

func main(){
	file, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var obj = make([]map[string] interface{} ,100)

	decodeJSON := json.NewDecoder(file)
	err = decodeJSON.Decode(&obj)

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(obj)
	// fmt.Println("%T",obj)
	// fmt.Println(obj[len(obj)-1])
	bigBooks := make([]Record,100)
	for index, value := range obj{
		// fmt.Println(index,value["author"],value["country"],value["year"])
		s := fmt.Sprintf("%f", value["year"])
		s = strconv.FormatFloat(value["year"].(float64), 'f', 0, 64)
		val , err := strconv.ParseInt(s, 10, 16)
		if err == nil {
			bigBooks = append(bigBooks, Record{author:value["author"].(string),year:uint16(val)})
			fmt.Println(index,value["author"],value["country"],value["year"],val,s)
		}
		
	}




}