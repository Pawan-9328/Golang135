package main

import (
	"fmt"

	
)

func main(){
 
   for i:= 0 ; i<10; i++ {
		  fmt.Println("Number is :", i)
	}

	counter := 0;
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 1 {
			break
		}
	}

	// slice 
	numbers := []int{1,2,3,4,5}
	for index , value := range numbers {
		 fmt.Printf("Index of numbers is %d, and value is %d\n :", index, value)
	}

	// range keywords make loop easy 
    data := "hello, world!"
	 for index, char := range data {
		 fmt.Printf("Index of Data %d , and value is %d\n", index, char)
	 }

}  