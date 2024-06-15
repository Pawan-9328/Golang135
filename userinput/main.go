package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hiii , What's your name ?")
	// var name string 

	// fmt.Scan(&name)
	// fmt.Println("hello, Mr. ", name)
    
	// ... read white space..... 
	 reader := bufio.NewReader(os.Stdin)
    name, _ :=	reader.ReadString('\n')

   fmt.Println("hello, Mr. ", name)

}