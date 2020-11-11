package main

// #include<stdio.h>
// void callC(){
// 	printf("Hello World!From C Programming\n");
// }
import "C"
import "fmt"
func main(){
	fmt.Println("Go statement")
	C.callC()
	fmt.Println("From Go program!")
}