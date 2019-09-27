package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}
	if data == 0 {
		fmt.Printf("value is %v\n", data)
	}
}
