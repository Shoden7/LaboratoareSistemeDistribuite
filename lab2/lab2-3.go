package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut"
		ch <- "pa"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}