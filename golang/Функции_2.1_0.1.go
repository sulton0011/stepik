// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

// Входные данные
// Вводится четыре числа.

// Выходные данные
// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input: 4 5 6 7

// Sample Output: 4

package main

import (
	"fmt"
)

func main () {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
    var a, min int = 0, 9999

	for i := 0; i < 4; i ++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
	}
	return min //print your code here
}
