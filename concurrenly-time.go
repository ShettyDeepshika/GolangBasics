package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int){
	time.Sleep(2*time.Second)
	fmt.Println(i*i)
}
func main(){
	start:=time.Now()
	for i:=1;i<20;i++{
		go calculateSquare(i)
	}
	//main go routine exists without waiting for all goroutines to complete
	time.Sleep(3*time.Second) 
	fmt.Println(time.Since(start))
}