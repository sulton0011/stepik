// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход
// число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
//
// Sample Input:
// 5
// 38 24 800 8 16

// Sample Output: 40

package main

import "fmt"

func main() {
	var a, sum int

	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		var value int
		fmt.Scan(&value)
		if value > 10 && value < 100 && value%8 == 0 {
			sum += value
		}
	}

	fmt.Println(sum)
}
