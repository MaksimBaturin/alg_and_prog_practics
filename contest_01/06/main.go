package main

import "fmt"

func main() {
	var N, FiveThousand, OneThousand, FiveHundred, TwoHundred, OneHundred int
	fmt.Scan(&N)
	FiveThousand = int(N / 5000)
	OneThousand = int((N - 5000*FiveThousand) / 1000)
	FiveHundred = int((N - 5000*FiveThousand - 1000*OneThousand) / 500)
	TwoHundred = int((N - 5000*FiveThousand - 1000*OneThousand - 500*FiveHundred) / 200)
	OneHundred = int((N - 5000*FiveThousand - 1000*OneThousand - 500*FiveHundred - 200*TwoHundred) / 100)
	fmt.Println(FiveThousand, OneThousand, FiveHundred, TwoHundred, OneHundred)
}
