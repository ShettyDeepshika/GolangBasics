package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	mypassword := "secret"
	originalpassword:=getMD5Hash(mypassword)
	var password string
	flag:=true
	temp:=1
	for flag&&temp<=3{
		fmt.Print("Enter your password: ")
		fmt.Scanln(&password)
		enteredPassword:=getMD5Hash(password)
		if originalpassword==enteredPassword{
			fmt.Println("Login successful")
			flag=false
		} else {
			fmt.Println("Invalid password")
			temp++
		}
	}
	if temp>3{
		fmt.Println("You have exceed the minimum number of login attempts.. Try again after 12 hours")
	}
	
}