package main

import "fmt"

func main() {

	day := 3

	switch day {
	case 1:
		{
			fmt.Println("Monday ")
		}
	case 2:
		{
			fmt.Println("Tuesday")
		}
	case 3:
		{
			fmt.Println("Wednesday")
		}
	case 4:
		{
			fmt.Println("Thursday")
		}
	default: {
		 println("unknown day")
	}
	}


	 month := "jan"

	 switch month{
	 case  "jan" , "feb" , "mar":
		   fmt.Println("winter")
	 
		case  "april" , "may" , "june":

		   fmt.Println("sprin")

		case  "july" , "august" , "september":
		   fmt.Println("raining")  
	 }

}
