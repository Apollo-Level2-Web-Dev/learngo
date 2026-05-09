package main

import "fmt"

func doSomething() {
	defer func() {
		fmt.Println("Deferred function ran")
		r := recover()
		if r != nil {
			fmt.Println("Recovered form", r)
		}
	}()

	panic("Something really bad happened")
}

func main() {
	doSomething()

	fmt.Println("Main completed normally")
}
