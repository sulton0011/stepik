// По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", 
// "n коровы", правильно склоняя слово "корова".

// Входные данные
// Дано число n (0<n<100).

// Выходные данные
// Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например,
// 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

// Sample Input:
// 10

// Sample Output:
// 10 korov

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