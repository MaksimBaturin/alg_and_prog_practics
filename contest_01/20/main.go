package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp, temp_n string
	var n1, n2 string
	fmt.Scan(&n1, &n2)
	for i := 0; i < len(string(n1)); i++ {
		temp_n = string(n1)[i : i+1]
		if strings.ContainsAny(n2, temp_n) == true {
			temp += "True"
		} else {
			temp += "False"
		}
	}
	if strings.Contains(temp, "False") {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
