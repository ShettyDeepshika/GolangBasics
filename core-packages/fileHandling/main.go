package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	//Join method that constructs paths in a portable way.
	path := filepath.Join("dir1","dir2/../dir3","text.txt")
	fmt.Println(path)
	//Dir method splits a path to the directory and the file.
	pathToDir := filepath.Dir(path)
	fmt.Println(pathToDir)
	//Base method returns last element of the path.
	basePath := filepath.Base(path)
	fmt.Println(basePath)
	//IsABs - checks if the path is absolute or not.
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs("/dir/file"))
	//Ext method gets the extension from a filename.
	fmt.Println(filepath.Ext(path))

	//
	fileInfo,err := os.Stat("/Users/M1080764/Desktop/GolangKodekloud/GolangBasics/core-packages/temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())	//size in bytes
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())

	//Reads the named file and returns the contents
	filePath := "/Users/M1080764/Desktop/GolangKodekloud/GolangBasics/core-packages/temp.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	
	//
	f,err := os.Open(filePath)
	if err !=nil {
		fmt.Println(err)
	}
	b := make([]byte ,4)
	for {
		n,err := f.Read(b)
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

	//OpenFile() method to append text to a file/create a file it it doen't exist
	file, err := os.OpenFile("/Users/M1080764/Desktop/GolangKodekloud/GolangBasics/core-packages/temp2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil{
		fmt.Println(err)
	}

	defer file.Close()
	
	//WriteSTring writes the contents of string rather than slice of bytes
	_, err = file.WriteString("Hope you had a good day!")
	if err != nil{
		fmt.Println(err)
	}
}