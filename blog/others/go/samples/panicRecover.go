package main

import "fmt"

func a(){
	fmt.Println("Inside func a");
	defer func (){
		if c := recover(); c != nil{
			fmt.Println("a: recovering from panic",c)
		}
	}()
	fmt.Println("calling b")
	b()
	fmt.Println("Exiting from a") // not executed
}

func b(){
	fmt.Println("Inside func b")
	panic("b: I am in Panic")
	fmt.Println("Exiting from b")// not executed
}

func main(){
	a();
	fmt.Println("Program terminated!")
}