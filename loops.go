package main

import "fmt"

func main(){
	i:=0
	for i<=5 {
		i++;
		fmt.Println(i);
	}
	if i<5 {
		fmt.Println("i < 5");
	} else {
		fmt.Println("i >= 5")
	}
	switch i {
	case 6 : fmt.Println("i>5")
	default: fmt.Println("default")
	}
}
