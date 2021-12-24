// По данному трехзначному числу определите, все ли его цифры различны.

// Формат входных данных
// На вход подается одно натуральное трехзначное число.

// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

// Sample Input 1:
// 237

// Sample Output 1:
// YES

// Sample Input 2:
// 117

// Sample Output 2:
// NO

package main

import "fmt"

func main() {
	var a int64
	fmt.Scan(&a)

	son1 := a % 10
	son2 := a / 10 % 10
	son3 := a / 100 % 10
	fmt.Println(son1, son2, son3)

	switch  {
		case son1 != son2 && son1 != son3 && son3 != son1 && son2 != son3:
			fmt.Println("YES")
		default:
			fmt.Println("NO")
	}
}