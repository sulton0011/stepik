// Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные
// Вводится натуральное число N, а затем N целых чисел последовательности.

// Выходные данные
// Выведите количество минимальных элементов последовательности.

// Sample Input: 3 21 11 4

// Sample Output: 1

package main

import "fmt"

func main () {
	var a, b, count int
	var min int = 999999999

	for fmt.Scan(&a); a > 0; a -- {
		fmt.Scan(&b)
		if b == min {
			count ++
		}
		if b < min{
			min = b
			count = 1
		}


	}
	fmt.Println(count)
}