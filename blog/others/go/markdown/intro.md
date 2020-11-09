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

The **log** package offers many functions for sending output to the syslog server of a UNIX
machine.

log.Printf() , log.Print() , log.Println() , log.Fatalf() , log.Fatalln() , log.Panic() , log.Panicln() , and log.Panicf()

There are various logging levels, including debug, info, notice, warning, err, crit, alert, and emerg, in reverse order of severity

##### Logging Facility

* A logging facility is like a category used for logging information.
* Its value can be one of auth, cron, daemon, kern, lpr, mail, mark, news, syslog, user, UUCP, local0, local1 etc... and is defined inside **/etc/syslog.conf** , **/etc/rsyslog.conf** or another appropriate file depending on the server process used for system logging on our UNIX machine.
* If it is not defined, then log messages we send to it might get ignored and therefore lost.

So In order to send the messages to the log file, there has to be an entry in configuration file

```
# server logging facility
server.* /var/log/server.log
```

The use of **log.Fatal()** terminates a Go program at the point where **log.Fatal()** was called

```
package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/custom.log"

func main(){
	f, err := os.OpenFile(LOGFILE, os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)

	if (err != nil){
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f,"customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Custom log entry")

}
```

```
cat /tmp/custom.log 
customLogLineNumber 2020/11/09 11:37:24 Hello there!
customLogLineNumber 2020/11/09 11:37:24 Custom log entry
```
```

iLog.SetFlags(log.LstdFlags | log.Lshortfile)
// Above code generates the below log

customLogLineNumber 2020/11/09 11:40:17 customLog.go:23: Hello there!
customLogLineNumber 2020/11/09 11:40:17 customLog.go:24: Custom log entry
```

#### Error Handling

```
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
```

#### Using docker

Dockerfile

```
FROM golang:alpine
RUN mkdir /files
COPY hello.go /files
WORKDIR /files
RUN go build -o /files/hw hw.go
ENTRYPOINT ["/files/hw"]
```

##### Building image

```
docker build -t hello_go:v1 .
```

##### running the application

```
docker run hello_go:v1
```
