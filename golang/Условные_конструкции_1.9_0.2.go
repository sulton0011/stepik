// Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.

// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.

// Sample Input:
// 1234

// Sample Output:
// 1

package main

import "fmt"

func main() {
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