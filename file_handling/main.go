package main

import (
	
 	"fmt"
// 	"io"
 	"io/ioutil"
// 	"os"
 )

func main() {

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file :", err)
	// 	return
	// }
	// defer file.Close()

	// content := "hello world by keshva"
	// // for writing return -  int , err
	// byte, errors := io.WriteString(file, content)
	// fmt.Println("byte written: ", byte)
	// if errors != nil {
	// 	fmt.Println("Error while writing file: ", errors)
	// }
	// fmt.Println("sucessfully created file")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("error while creating file", err)
	// 	return
	// }
	// defer file.Close()

	// // read data using buffer
	// // created a buffer to read the file content

	// buffer := make([]byte, 1024)
	// // read the file content into the buffer
	// for {
	// 	n, err := file.Read(buffer)
	// 	 // EOF - end of file 
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error while reading file", err)
	// 		return
	// 	}

	// 	// process the read content
	// 	fmt.Println(string(buffer[:n]))
	// }
      
	// Read the entire file into a byte slice 
   
	// also used os 
	content , err := ioutil.ReadFile("example.txt")
	if err != nil {
		  fmt.Println("error while reading file ", err)
		  return 
	}
   fmt.Println(string(content))

}
