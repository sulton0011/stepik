package main

import "fmt"

func main () {
	var a int32
	fmt.Scan(&a)

	switch  {
		case a < 10:
			a = a % 10
		case a < 100:
			a = a / 10 % 10
		case a < 1000:
			a = a / 100 % 10
		case a < 10000:
			a = a / 1000 % 10
		case a < 100000:
			a = a / 10000 % 10
	}

	fmt.Println(a)
}