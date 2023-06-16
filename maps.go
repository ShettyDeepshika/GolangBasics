package main

import "fmt"

func main(){
	presAge:=make(map[string] int)
	presAge["PersonA"] = 50
	fmt.Println(presAge["PersonA"])
	presAge["PersonB"] = 43
	fmt.Println(len(presAge))
	
	// We can delete by passing the key to delete
	
	delete(presAge, "PersonB")
	fmt.Println(len(presAge))
}
//////////////////////
// package main
// import "fmt"
// func main(){
// 	mapEx:=make(map[string]int)
// 	mapEx["A"]=65
// 	mapEx["F"]=70
// 	mapEx["K"]=75
// 	delete(mapEx,"F")
// 	// for i,value:=range mapEx{
//     	fmt.Println(mapEx)
// 	// }
// }
