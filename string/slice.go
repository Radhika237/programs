package main

import (
	"fmt"
)

func difference(slice1 []string, slice2 []string) []string {

	strSlice := append(slice1, slice2...)

	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func main() {
	slice1 := []string{"Red", "Black", "White"}
	slice2 := []string{"Black", "Yellow", "Orange"}

	fmt.Printf("%+v\n", difference(slice1, slice2))
}
