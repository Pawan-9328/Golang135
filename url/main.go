package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning url")
	myURL := "https://www.google.com/"
	fmt.Printf("Type of URL : %T\n", myURL)

	parseURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	fmt.Printf("Type of URL : %T\n", parseURL)

	fmt.Println("Schema : ", parseURL.Scheme)
	fmt.Println("Host : ", parseURL.Host)
	fmt.Println("Path : ", parseURL.Path)	
	fmt.Println("RawQuery ", parseURL.RawQuery)

	//.......modifying URL components......

	 parseURL.Path = "/newPath"
	 parseURL.RawQuery = "username=iamkeshav"

	 //.....constructing a URL string from a URL object......
   newUrl  := parseURL.String() 
	fmt.Println("new URL: ", newUrl)


}
