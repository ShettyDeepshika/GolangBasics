package main

import "fmt"

func main(){
	x:=0;
	changeXByValue(x);
	fmt.Println("X = ",x);
	changeXByAddress(&x);
	fmt.Println("X = ",x)
	fmt.Println("Address of X = ",&x)
}
func changeXByValue(x int){
	x=1;
}
func changeXByAddress(x *int){
	*x=2;
}