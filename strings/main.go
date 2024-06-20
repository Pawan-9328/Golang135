package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple, orange,banana,Prince"
	parts := strings.Split(data, ",") 
	fmt.Println(parts)
   
	// how many time two repeat in this function
   str := "one two three four two two five"
	count := strings.Count(str,"two")
	fmt.Println("count: ", count) 
   
	// trimmed white space  
	str = "                    Hello, Go!              "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed: ", trimmed)

	str1 := "Keshav"
	str2 := "Gola"

	result := strings.Join([]string{str1,str2}, " ")
	fmt.Println("result: ", result)
}
