// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

// Sample Input: ihgewlqlkot
// Sample Output: hello

package main

import (
	. "fmt"
)

func main () {
	var s, s1 string

	Scan(&s)

	for j, i := range s {
		if j % 2 != 0 {
			s1 = s1 + string(i)
		}
	}

	Println( s1)
}