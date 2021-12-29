// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

// Sample Input: zaabcbd

// Sample Output: zcd

package main

import (
	. "fmt"
	"strings"
)

func main () {
	var s string

	Scan(&s)

	for _, i := range s {
		if strings.Count(s, string(i)) != 1 {
			s = strings.ReplaceAll(s, string(i), "")
		}
	}

	Println(s)
}