package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main(){
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO | syslog.LOG_MAIL, programName )

	if err != nil{
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}

	log.Println("Writing to the log file: Through go program")
	fmt.Println("check the log message")

}