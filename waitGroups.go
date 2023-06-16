package main

import (
	"fmt"
	"time"
	"sync"
)

func calculateSquare(i int, wg *sync.WaitGroup){
	fmt.Println(i*i)
	wg.Done()
}
func main(){
	var wg sync.WaitGroup
	start:=time.Now()
	wg.Add(10)
	for i:=0;i<10;i++{
		go calculateSquare(i,&wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}