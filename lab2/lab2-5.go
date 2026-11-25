package main

import "fmt"

func send(ch chan<- string) {
	ch <- "salut"
}

func receive(ch <-chan string, out chan<- string) {
	out <- <-ch
}

func main() {
	ch := make(chan string)
	out := make(chan string)

	go send(ch)
	go receive(ch, out)

	fmt.Println(<-out)
}