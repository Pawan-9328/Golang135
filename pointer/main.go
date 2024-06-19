package main

import "fmt"

func modifyValueRef(num *int) {
	//change  value
	*num = *num * 20

}

func main() {

	// var num int
	num := "Keshav"

	// var ptr *int
	ptr := &num

	//fmt.Println("Num has value: ", num)
	fmt.Println("Pointer contains: ", ptr)
	fmt.Println("Data contains through pointer :", *ptr)

	var pointer *int // default pointer == nil
	if pointer == nil {
		fmt.Println()
	}
	value := 10
	modifyValueRef(&value)
	fmt.Println("Value contains :", value)
}
