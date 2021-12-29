// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка)
// между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

// Выходные данные

// Вывести строку, которая получится после добавления символов '*'.

// Sample Input: LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
// Sample Output: L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O

// Sample Input: LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
// Sample Output: L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O

package main

import (
	"fmt"
	"strings"
)

func main () {
	var s string

	fmt.Scan(&s)

	fmt.Println(strings.ReplaceAll(s, "", "*")[1:len(s) * 2])
}