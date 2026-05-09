package main

import (
	"fmt"
	"log"
)

func doSomething() {
	defer func() {
		fmt.Println("Deferred function ran")
		r := recover()
		if r != nil {
			fmt.Println("Recovered form", r)
		}
	}()

	panic("Something really bad happened") // missile
}

func doAnotherThing() {
	defer func() {
		fmt.Println("Deferred function ran")
	}()

	log.Fatal("something very big happened")

}

func main() {
	// doSomething()
	doAnotherThing()

	fmt.Println("Main completed normally")
}
