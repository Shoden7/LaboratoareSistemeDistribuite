package main

import "fmt"

func main() {
	ch := make(chan string)
	done := make(chan bool)

	go func() {
		fmt.Println(<-ch)
		close(done) 
	}()

	ch <- "salut"
	<-done
}