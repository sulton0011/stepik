// Самое большое число, кратное 7

// Найдите самое большее число на отрезке от a до b, кратное 7 .

// Входные данные
// Вводится два целых числа a и b (a≤b).

// Выходные данные
// Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или
// выведите "NO" - если таковых нет.

// Sample Input: 100 500

// Sample Output: 497

package main

import "fmt"

func main() {
	var a, c int64
	var i int = 1

	fmt.Scan(&a, &c)

	if c == 0 {
		fmt.Println(c)
		i = 0
	} else {
		for ; c >= a; c-- {
			if c%7 == 0 && c != 0 {
				fmt.Println(c)
				i = 0
				break
			}
		}
	}

	if i != 0{
		fmt.Println("NO")
	}
}
