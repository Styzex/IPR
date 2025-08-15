package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Resolver = "https://checkip.amazonaws.com/"

func ResolveIP() (string, error) {
	Response, Error := http.Get(Resolver)

	if Error != nil {
		return "",  Error
	}

	defer Response.Body.Close()
	Content, Error := ioutil.ReadAll(Response.Body)
	if Error != nil {
		return "", Error
	}
	
	return string(Content), nil
}

func main() {
	IP, Error := ResolveIP()
	
	if Error != nil {
		fmt.Println("Error:", Error)
	} else {
		fmt.Println("Your public IP is: " + IP)
	}
}