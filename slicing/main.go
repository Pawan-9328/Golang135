package main

import (
	"fmt"


)

func main() {

 	//numbers := []int{1, 2, 3, 4, 5}
 	//numbers = append(numbers, 3,4,5,6,7,8,9,10)
// 	fmt.Println("Number : ", numbers)
// 	fmt.Printf("Numbers has data type : %T\n", numbers)
// 	fmt.Println("Length : ", len(numbers))

//  // slice intialize 
//  name := []string{}
//  fmt.Println("name : ", name)

numbers := make([]int,3,5)



numbers = append(numbers, 2)
numbers = append(numbers, 4)
numbers = append(numbers, 67)
numbers = append(numbers, 87)
numbers = append(numbers, 90)
numbers = append(numbers, 70)
numbers = append(numbers, 23)
numbers = append(numbers, 15)

 
fmt.Println("Slice:", numbers)
fmt.Println("Length:", len(numbers))
fmt.Println("Capacity:", cap(numbers))

}