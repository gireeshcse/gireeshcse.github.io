package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main(){
	// programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_ALERT | syslog.LOG_MAIL, "some program" )

	if err != nil{
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
	fmt.Println("check the log message")

}