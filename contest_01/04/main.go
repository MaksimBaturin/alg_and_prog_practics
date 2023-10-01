package main

import "fmt"

func main() {
	var a int32
	var b int32
	fmt.Scan(&a, &b)
	var c int32 = a * b
	fmt.Println(c)
}
