package main

import (
	"fmt"
)

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(password string) error {
	if password != "1234" {
		return &CustomError{
			message: "Password do not match",
			code:    401,
		}

		// return errors.New("password do match")

	}
	return nil
}

func main() {

	err := login("2345")
	if err != nil {

		// customErr, ok := err.(*CustomError)

		// if !ok {
		// 	fmt.Println(customErr)
		// } else {

		// 	fmt.Println(customErr)
		// 	fmt.Println(customErr.code)
		// }

		if customError, ok := err.(*CustomError); ok {
			fmt.Println(customError.code)
		}

		// fmt.Println("Error", err, "Code", err.code)
	}

	users := map[int]string{
		1: "Mezba",
		2: "Mir",
		3: "Firoz",
	}

	name, ok := users[2] // ""

	fmt.Println(name)
	fmt.Println(ok)

	if ok {

	}

	fmt.Println("main ends")

}
