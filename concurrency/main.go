package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var start = time.Now()

	// wg.Add(1) // 1
	// go uploadFile()

	// wg.Add(1) // 2
	// go saveToDB()
	// wg.Add(1) // 3
	// go sendEmail()

	wg.Go(uploadFile)
	wg.Go(saveToDB)
	wg.Go(sendEmail)

	wg.Wait() // waiting .... until counter is 0

	fmt.Println("all tasks completed")
	fmt.Println("time taken", time.Since(start))
}

func uploadFile() {
	// defer wg.Done()
	fmt.Println("uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")

	// wg.Add(-1) // 2
}
func saveToDB() {
	// defer wg.Done()
	fmt.Println("saving to db...")
	time.Sleep(1 * time.Second)
	fmt.Println("saved to db!")

	// wg.Add(-1) // 1
}
func sendEmail() {
	// defer wg.Done()
	fmt.Println("sending email...")
	time.Sleep(2 * time.Second)
	fmt.Println("email sent!")

	// wg.Add(-1) // 0
}
