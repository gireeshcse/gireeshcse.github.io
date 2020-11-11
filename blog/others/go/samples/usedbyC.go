package main

import "C"
import "fmt"

//export HelloGo
func HelloGo(){
	fmt.Println("Hello from GO!")
}

//export Multiply
func Multiply(a,b int) int{
	return a*b
}

func main(){

}