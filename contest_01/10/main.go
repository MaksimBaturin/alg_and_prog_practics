package main

import "fmt"

func main() {
	var swim, wings, neck string
	fmt.Scan(&swim, &wings, &neck)
	if swim == "Да" {
		if wings == "Да" {
			if neck == "Да" {
				fmt.Println("Утка")
			} else {
				fmt.Println("Пингвин")
			}
		} else {
			if neck == "Да" {
				fmt.Println("Плезиозавры")
			} else {
				fmt.Println("Дельфин")
			}
		}
	} else {
		if wings == "Да" {
			if neck == "Да" {
				fmt.Println("Страус")
			} else {
				fmt.Println("Курица")
			}
		} else {
			if neck == "Да" {
				fmt.Println("Жираф")
			} else {
				fmt.Println("Кот")
			}

		}
	}
}
