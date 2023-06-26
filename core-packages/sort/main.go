package main

import (
	"fmt"
	"sort"
)

func main(){
	vars := []int{5, 6, 1, 8,10}
	str := []string{"a","n","cd","edf"}
	sort.Ints(vars)
	sort.Strings(str)
	fmt.Println(vars)
	fmt.Println(str)
}