package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Learning web services ")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}
	defer res.Body.Close() 
	fmt.Printf("Type of response: %T\n", res)
	//fmt.Println("response: ", res)
         
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		 fmt.Println("Error reading response", err)
		 return 
	}
  // convert into string 
	fmt.Println("response: ", string(data));
}
