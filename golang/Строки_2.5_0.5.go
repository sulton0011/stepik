// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля
// должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается
// строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

// Sample Input: fdsghdfgjsdDD1
// Sample Output: Ok

package main

import (
	. "fmt"
	"unicode"
	"unicode/utf8"
)

func main () {
	var s string
	var count int

	Scan(&s)

	for _, i := range s {
		if unicode.Is(unicode.Latin, rune(i)) == true || unicode.IsDigit(rune(i)) == true {
			count ++
		}
	}

	if count == utf8.RuneCountInString(s) && utf8.RuneCountInString(s) >= 5{
		Println("Ok")
	}else {
		Println("Wrong password")
	}
}