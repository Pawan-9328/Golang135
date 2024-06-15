package main 

import "fmt"

//...without parameterized function........
func simpleFunction(){
	 fmt.Println("simple function ")
}

//fucntion with parameter and return type 
func add(a, b  int) int {
	  return a+b; 
}

//.......both function are same.........
// func add2(x int , y int){
// 	return x+y; 
// }

//...function with named return variables 
func Mines(n int , m int) (result int){
	 result = n+m; 
	 return result
		
}

func Mutiplication(r , u int) (result int){
	 result = r*u
	 return 
}


func main(){
	 fmt.Println("We are using functuon ")
	 simpleFunction()

	 ans := add(3,5);
	 fmt.Println("sum of two numbers is: ",ans); 

}
