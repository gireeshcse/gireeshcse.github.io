### Go

* Go is a modern, generic-purpose, open source programming language
* Designed to to build reliable,robust, and efficient software
* Go code is portable, supports garbage collection, procedural ,concurrent and distributed programming.
* Go uses static linking by default-- binary files does not need any libraries or dependencies for execution once compiled.
* In go, we can either use a Go package or we do not include it for later usage.
* The underscore character, which is called blank identifier, is the Go way of discarding a value

#### Cons

* No direct support of OOP, but we can use composition in Go to mimic inheritance.
* C is still faster than other programming language for systems programming.

#### godoc utility

```
go doc fmt.Printf
go doc fmt
```

#### Sample Program

hello.go
```
package main
import (
    "fmt"
)
func main(){
    fmt.Println("Hello World!")
}
```

compiling a statically linked executable file
```
go build hello.go
```
To Execute
```
./hello
go run hello.go
```

Go requires the use of semicolons as statement terminators in many contexts, and the compiler automatically inserts the required semicolons when it thinks that they are necessary. Therefore, putting the opening curly brace ( { ) in its own line will make the Go compiler insert a semicolon at the end of the previous line ( func main() for example), which will cause of the error message if it contains function defination.

#### Downloading external packages

    go get -v github.com/gireeshcse/go/usefulPackage

downloads to the following 

    ~/go/src/github.com/gireeshcse/go/usefulPackage

go get also compiles and these can be found at

    ~/go/pkg/linux_amd64/gireeshcse/go/usefulPackage.a

Delete intermediate files of a downloaded go package

    go clean -i -v -x github.com/gireeshcse/go/usefulPackage

Entire package can be deleted using

    go clean -i -v -x github.com/gireeshcse/go/usefulPackage
    rm -f ~/go/pkg/linux_amd64/gireeshcse/go/usefulPackage

#### Printing Output

* fmt.Println
* fmt.Print
* fmt.Printf
* fmt.Sprintf
* fmt.Sprint
* fmt.Sprintln
* fmt.fprintf
* fmt.fprint
* fmt.fprintln

```
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Give one more argument"
	}else{
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout,"\n")
}
```

```
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">" , scanner.Text())
	}
}
```

```
n, _ := strconv.ParseFloat(arguments[i], 64)
io.WriteString(os.Stderr, myString)
go run stdERR.go >/tmp/output 2>&1
```

#### Logging

log.Printf() , log.Print() , log.Println() , log.Fatalf() , log.Fatalln() , log.Panic() , log.Panicln() , and log.Panicf()

There are various logging levels, including debug, info, notice, warning, err, crit, alert, and emerg, in reverse order of severity

/etc/syslog.conf , /etc/rsyslog.conf
syslogd and rsyslogd