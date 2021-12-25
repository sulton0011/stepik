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
		if value%8 == 0 {
			for value > 0 {
				sum += value % 10
				value /= 10
			}
		} else {
			var count, sum1 int
			for value > 0 {
				sum1 += value % 10
				if (value%10)%8 == 0 {
					count++
				}
				value /= 10
			}
			if count != 0 {
				sum += sum1
			}
		}
	}

	fmt.Println(sum)
}
