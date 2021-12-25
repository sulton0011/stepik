// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

// Входные данные
// Вводится натуральное число.

// Выходные данные
// Выведите ответ на задачу.

// Sample Input: 50

// Sample Output: 1 2 4 8 16 32

package main

import "fmt"

func main() {
	var (
		a     int64
		d     int64
		b     int64 = 1
		c     int64 = 1
		count int   = 2
	)

	fmt.Scan(&a)

	if a >= 2 {
		for i := int64(2); d < a; i++ {
			d = b + c
			c = b
			b = d
			count++
		}
	}

	if d == a {
		fmt.Println(count)
	} else {
		fmt.Println(-1)
	}
}