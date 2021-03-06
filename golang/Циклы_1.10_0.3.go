// Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
// которые равны ее наибольшему элементу.
//
// Формат входных данных
// Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит,
// а служит как признак ее окончания).
//
// Формат выходных данных
// Выведите ответ на задачу.
//
// Sample Input:
// 1
// 3
// 3
// 1
// 0
//
// Sample Output:
// 2

package main

import (
	"fmt"
)

func main() {
	var arr = []int{}
	max := 0
	a := 1
	for a != 0 {
		fmt.Scan(&a)
		if a > max {
			max = a
		}
		if a != 0 {
			arr = append(arr, a)
		}
	}

	var count int
	for _, i := range arr {
		if i == max {
			count++
		}
	}
	fmt.Println(count)
}
