package main

import (
	log "github.com/sirupsen/logrus"
)

/* Log levels
Trace
Debug
Info
Warn
Error
Fatal
Panic
*/
func main(){
	// file, err := os.OpenFile("/Users/M1080764/Desktop/GolangKodekloud/GolangBasics/core-packages/temp2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// log.SetOutput(file)
	log.Println("Hello World")
	log.Trace("Hello World")	//output--- time="2023-06-22T11:45:54+05:30" level=info msg="Hello World"
	// log.SetLevel(log.DebugLevel)
	log.Debug("Hello World")	//output--- time="2023-06-22T11:45:54+05:30" level=debug msg="Hello World"
	log.Warn("Hello World")		//output--- time="2023-06-22T11:47:09+05:30" level=warning msg="Hello World"
	log.Error("Hello World")	//output--- time="2023-06-22T11:47:09+05:30" level=error msg="Hello World"
	// log.GetLevel()
	log.Fatal("Hello World")	// output ----- time="2023-06-22T11:48:09+05:30" level=fatal msg="Hello World" exit status 1
	log.Panic("Hello World")	//output ---- panic: (*logrus.Entry)

}