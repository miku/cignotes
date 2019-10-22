package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	salutation := "hello"
	// wg.Add(1)
	go func() {
		//defer wg.Done()
		salutation = "welcome"
	}()
	// salutation = "3"
	// wg.Wait()
	fmt.Println(salutation)
}
