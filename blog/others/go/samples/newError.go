package main

import (
	"fmt"
	"errors"
	"os"
)

func returnError(a , b int) error {
	if ( a == b){
		err := errors.New("Error in returnError function")
		return err
	}

	return nil;

}

func main(){
	err := returnError(10, 9)
	if err == nil {
		fmt.Println("Normal execution")
	}else{
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("Normal execution")
	}else{
		fmt.Println(err)
		// converts error to string
		fmt.Println(err.Error())
		panic(err) // built in function stops the execution of a program
		os.Exit(10)
	}
}