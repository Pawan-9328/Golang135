package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"strings"

	//"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json : "id"`
	Title     string `json: "title`
	Completed bool   `json "completed"`
}

func performGetRequest() {
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()
	// also you used directly 200
	if res.StatusCode != http.StatusOK {
		fmt.Println("Errro handling ", res.Status)
	}
	// ...........generally can't is used this type methods bcz this return only json then unmarshell then store in todos............

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading ", err)
	// 	return
	// }
	// fmt.Println("Data : ", string(data))

	var todo Todo
	// NewDecoder - data decode in normal objects and save in Todo
	json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decodein :", err)
		return
	}
	fmt.Println("Todo: ", todo)

	fmt.Println("completed response : ", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		//UserId: 23,
		Title:     "keshav",
		Completed: true}

	// convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshling", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myURL := ""
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
	}
	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response : ", string(data))

	 fmt.Println("response status :", res.Status)
}

func main() {
	fmt.Println("CRUD")
	performGetRequest()
	performPostRequest()

}

