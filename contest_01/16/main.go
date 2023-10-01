package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	Numb := make([]float64, n)

	for i := 0; i < len(Numb); i++ {
		fmt.Scan(&Numb[i])
	}

	fmt.Print(Numb[0])

	for i := 1; i < (len(Numb) - 1); i++ {
		fmt.Print(" ", ((Numb[i-1] + Numb[i] + Numb[i+1]) / 3))
	}

	fmt.Print(" ", Numb[n-1])
}
