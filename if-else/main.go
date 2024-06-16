package main

import "fmt"

func main() {

	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		println("x is smaller than 5")
	}

	z := 21
	if z > 11 {
		fmt.Println("z is greater than 21")
	} else if z > 5 {
		fmt.Println("z greater than 5 but smaller than 5")
	} else {
		fmt.Println("z is smaller than 5")
	}


	y := 10

	if x < 5 && (y>5 || z<43) {
		   fmt.Println("how are you ? ")
	} else{
		 fmt.Println("Learn Programming with Hello World")
	}

}
