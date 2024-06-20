package main

import (
	"fmt"
	"time"
)

func main() {
	//"dd-mm-yyyy"
	//"yyyy/mm/dd"

	currentTime := time.Now()
	fmt.Println("Current time : ", currentTime)
	fmt.Printf("Type of currentTime %T \n", currentTime)

	//formatted := currentTime.Format("02-01-2006, 15:04:05")
	formatted := currentTime.Format("2006/01/02, 3:04 PM")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str,dateStr)
	fmt.Println("Formatted time: ", formatted_time)

	// add 1 more dqy in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date time: ", new_date) 

}
