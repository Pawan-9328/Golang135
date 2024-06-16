package main

import "fmt"

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0 , fmt.Errorf("Dominator must be not zero")
// 	}
// 	// dimision possiable 
// 	return a/b, nil
// }


// now we return string 

func divide(a, b float64) (float64, string ) {
	if b == 0 {
		return 0 , "Dominator must be not zero"
	}
	// dimision possiable 
	return a/b, "nil"
}

func main() {
	fmt.Println("started Error handling in the function")
	// in this case  Golang return infinite
	// underscore _ used for ignore the errors 
	ans, _ := divide(10, 2)
	//..........ex - now used error  / but most of used underscore...........
	// ans, err := divide(10,2)
	// if err != nil {
	// 	 fmt.Println("Error handling")
	// }
	fmt.Println("Division of two numbers is ", ans)
}
