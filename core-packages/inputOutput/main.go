package main

import (
	"fmt"
	"strings"
	// "io"
	// "os"
	// "log"
)


func main(){
	r := strings.NewReader("learning is fun")

	buf := make([]byte, 4)

	for{
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]),err)
		if err != nil {
			fmt.Println("breaking out")
			break
		}
	}

	//Copy - copies from src to dest until eitehr EOF is reached on src or an error occurs.
	// _,err := io.Copy(os.Stdout, r)
	// if err != nil{
	// 	log.Fatal(err)
	// }
}