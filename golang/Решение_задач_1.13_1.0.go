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
	var n int

	fmt.Scan(&n)

	if n >= 5 && n <= 20 {
		fmt.Println(n, "korov")
	}else if n % 10 == 2 || n % 10 == 3 || n % 10 == 4{
		fmt.Println(n, "korovy")
	}else if n % 10 == 1{
		fmt.Println(n, "korova")
	}else {			
		fmt.Println(n, "korov")
	}
}