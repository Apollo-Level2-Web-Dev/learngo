package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var ch = make(chan string) // un-buffered channel

	wg.Add(1)
	go uploadFile(ch)

	wg.Wait()

	// fileUrl := <-ch // blocking ...

	// fmt.Println("File Url", fileUrl)

}

func uploadFile(c chan string) {
	defer wg.Add(-1)

	fmt.Println("uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")
	fileUrl := "https://s3.3424234.png"
	c <- fileUrl

}
