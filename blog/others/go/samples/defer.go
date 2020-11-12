package main
import "fmt"

func func1(){
	for i := 3; i > 0; i--{
		defer fmt.Print(i) // 1 2 3
	}
	fmt.Println()
}

func func2(){
	for i := 3; i > 0; i--{
		defer func() {
			fmt.Print(i) // 0 0 0
		}()
	}
	fmt.Println()
}

func func3(){
	for i := 3; i > 0; i--{
		defer func(n int) {
			fmt.Print(n) // 1 2 3
		}(i)
	}
	fmt.Println()
}

func main(){
	fmt.Println("Defer program")
	func1()
	func2()
	func3()
}