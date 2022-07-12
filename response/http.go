package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com/apple")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("status code :", resp.StatusCode)
	fmt.Println(" err", http.StatusText(resp.StatusCode))
	fmt.Errorf(err.Error())
}
