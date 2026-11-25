package main

import "fmt"

func main() {
	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println("se proceseaza")
	}()

	<-done
	fmt.Println("main: goroutina a terminat")
}