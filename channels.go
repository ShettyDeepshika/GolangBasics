package main
import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup
func main(){
	ch:=make(chan string)
	wg.Add(2)
	go sell(ch)
	go buy(ch)
	// time.Sleep(5*time.Second)
	wg.Wait()
}
func sell(ch chan string){
	ch<-"Furniture"
	fmt.Println("Sent data to the channel....")
	wg.Done()
}
func buy(ch chan string){
	fmt.Println("Waiting to recieve data......")
	val:=<-ch
	time.Sleep(1*time.Second)
	fmt.Println("Received data from the channel: ", val)
	wg.Done()
}