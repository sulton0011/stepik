// Из натурального числа удалить заданную цифру.

// Входные данные
// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные
// Вывести число без заданных цифр.

// Sample Input:
// 38012732
// 3

// Sample Output:
// 801272

package main

import (
	"fmt"
	"strings"
)

func main () {
	var a, b string

	fmt.Scan(&a, &b)

	a = strings.ReplaceAll(a, b, "")

	fmt.Println(a)
}