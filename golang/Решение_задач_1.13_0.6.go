// Количество нулей
// По данным числам, определите количество чисел, которые равны нулю.  

// Входные данные 
// Вводится натуральное число N, а затем N чисел.

// Выходные данные 
// Выведите количество чисел, которые равны нулю.

// Sample Input: 5 1 8 100 0 12

// Sample Output: 1

package main

import "fmt"

func main () {
	var a, b, count int

	for fmt.Scan(&a); a > 0; a -- {
		fmt.Scan(&b)
		if b == 0 {
			count ++
		}
	}
	fmt.Println(count)
}