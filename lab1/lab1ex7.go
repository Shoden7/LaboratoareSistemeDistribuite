package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("cuvinte.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var cuv []string
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		cuv = append(cuv, sc.Text())
	}

	cnt := make(map[string]int)
	for _, w := range cuv {
		cnt[w]++
	}

	for w, c := range cnt {
		fmt.Println(w, c)
	}
}