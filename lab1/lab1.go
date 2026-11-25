package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//ex1
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	x, err := strconv.Atoi(text)
	var ok bool
	ok = false
	if err != nil {
		fmt.Println("Eroare la conversie")
	}

	if x < 2 {
		fmt.Println("Nu e prim")
	} else {

	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			fmt.Println("Nu e prim")
			ok = true
		}
	}


	if ok == false {
	fmt.Println("Prim")
	}
	}


	//ex2
	a := [...]int{5, 2, 9, 1, 7, 3, 14, 10}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	fmt.Println(a)

	//ex3
	b := [...]int{5, 2, 9, 1, 7, 3, 14, 10}
	var slicepar []int
	for i := 0; i < len(b); i++ {
		if b[i]%2 == 0 {
			slicepar = append(slicepar, b[i])
		}
	}
	fmt.Println(slicepar)


	//ex4
	vf := []float32{1.5, 2.3, 3.7, 4.1, 5.9}
	var i int
	fmt.Scan(&i)
	if i < 0 || i >= len(vf) {
		fmt.Println("Index invalid")
	}
	vf = append(vf[:i], vf[i+1:]...)
	fmt.Println(vf)

	//ex5
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)

	//ex6

	m := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George":  178,
		"Ana":    165,
	}
	for len(m) > 0 {
		Mnume, Minaltime := "", -1
		for nume, inaltime := range m {
			if inaltime > Minaltime {
				Mnume = nume
				Minaltime = inaltime
			}
		}
		fmt.Printf("%s: %d\n", Mnume, Minaltime)
		delete(m, Mnume)
	}

	//ex8
	angajati := map[string]int{
		"Ana Popescu":   1,
		"Mihai Ionescu": 2,
		"Bogdan Radu":   1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":  1,
	}

	departamente := make(map[int][]string)
	for nume, id := range angajati {
		departamente[id] = append(departamente[id], nume)
	}

	fmt.Println(departamente)

	//ex9
	c := [...]int{2, 7, 11, 8, 4, 9}
	var suma int
	fmt.Scan(&suma)

	seen := make(map[int]int)
	for i, v := range c {
		if j, ok := seen[suma-v]; ok {
			fmt.Println(j, i)
			
		}
		seen[v] = i
	}

}	