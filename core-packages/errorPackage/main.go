package main

import (
	// "errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		// return errors.New("only odd numbers are allowed")
		return fmt.Errorf("only odd numbers are allowed, got : %d",i)
	}
	return nil
}

func checkError(err error){
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Opeartion successful")
}
func main(){
	err := process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}