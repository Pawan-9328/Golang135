package main

import "fmt"

func main() {

	// names <-> grades

	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 34
	studentGrades["Alice"] = 98
	studentGrades["Bob"] = 21
	studentGrades["Charli"] = 88

	fmt.Println("marks of bob: ", studentGrades["Bob"])
	studentGrades["Bob"] = 100;
	 
	// checking if a key exits 
	grades, exits :=  studentGrades["Bob"]
	//fmt.Println("Marks of keshav :", studentGrades["keshav"])
	fmt.Println("Grades of Bob : ", grades)
	fmt.Println("Bob exits : ", exits)
	
  // range always privide index and value 
	for index , value  := range studentGrades {
		   fmt.Println("key is %s and marks is %d \n ", index,value)
	}
  
        person := map[string]int {
			      
			    "Alice" : 90,
				 "Bob": 56,
				 "Num" : 89,
		  }
	   for index , value := range person {
		   fmt.Println("key is %s and marks is %d \n ", index,value)
		
		}
}