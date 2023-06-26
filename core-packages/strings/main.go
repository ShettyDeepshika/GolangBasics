//This program demonstrate the use of "strings" package in Go
package main

import (
	"fmt"
	"strings"
)
func main(){
	learning:="learning standard library in Go"
	fun:="learning"

	//"Contains" in strings package used to search whether a particluar text/string/character is present in the given string.
	result:=strings.Contains(learning,fun)
	fmt.Println(result)

	//"Count" in strings package counts the number of times a word occur in a given sentence
	count:=strings.Count(learning,fun)
	fmt.Println(count)

	//"ReplaceAll" in strings package replace old text with a new text
	finalString:=strings.ReplaceAll(learning,fun,"teaching")
	fmt.Println(finalString)
}