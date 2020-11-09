package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)


func main(){
	if (len(os.Args) == 1){
		fmt.Println("Give more than 1 or more floats")
		os.Exit(1)
	}

	var err error = errors.New("error!")
	arguments := os.Args
	k := 1
	floatValues := 0
	var n float64
	var initValue float64
	for(k < len(arguments)){
		n, err = strconv.ParseFloat(arguments[k],64)
		if(err == nil){
			floatValues++
			initValue = n
		}
		k++
	}
	if(floatValues == 0){
		fmt.Println("Arguments or not float")
		return
	}

	min, max  :=  initValue, initValue

	for i := 1; i < len(arguments); i++ {
		n, err = strconv.ParseFloat(arguments[i],64)
		if err == nil{
			if (n < min){
				min = n
			}
			if (n > max){
				max = n
			}
		}
	}

	fmt.Println("Min: ",min)
	fmt.Println("Max: ", max)

}