package main

import "fmt"

func main() {
	fmt.Println("We are learning array in golang ")


	var name[5]string 
   
	// indexing start at 0 1 2 3 4 
	name[0] = "Keshav"
	name[1] = "prince"

// 	fmt.Println("Names of Person is: ", name)

// 	var numbers = [5]int{1,2,3,4,5}
// 	fmt.Println("Numbers is :", numbers)
//   // lenght finding in array 
// 	fmt.Println("lengths of numbers array is: ", len(numbers))
// 	//array element  accesing 
	//fmt.Println("value of name at 2 index is: ", name[2])

  // if any intailizing any space that time  golang set data  default value 
  // in this case set int default value which is zero.... 
   var price[5]string 
	fmt.Println("Price is: ",price)
	fmt.Printf("Price array is %q\n ",price)


}