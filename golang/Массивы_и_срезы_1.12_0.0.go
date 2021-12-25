// Внутри функции main (объявлять функцию не нужно) необходимо написать программу:

// На первом этапе на стандартный ввод подается 10 целых положительных чисел, которые 
// должны быть записаны в порядке ввода в массив из 10 элементов. Тип чисел, входящих в 
// массив, должен соответствовать минимально возможному целому беззнаковому числу.
// Имя массива который вы должны сами создать workArray (условие обязательное). 
// Для чтения из стандартного ввода уже импортирован пакет fmt.

// На втором этапе на стандартный ввод подаются еще 3 пары чисел - индексы элементов 
// этого массива, которые требуется поменять местами (если такая пара чисел 3 и 7, 
// значит в массиве элемент с 3 индексом нужно поменять местами с элементом, индекс 
// которого 7).

// Элементы полученного массива должны быть выведены через пробел на стандартный вывод.
// Далее автоматически будет проведена проверка используемых типов, результат которой
// будет добавлен к вашему ответу.

// Использование массива - обязательное условие!

// Sample Input:
// 99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8

// Sample Output:
// 99 137 151 187 117 71 20 187 93 67 type ok

package main 
import (
	"fmt"
)

func main () {
	var workArray [10]uint8
	for i := 0; i < 10; i ++ {
		var s uint8
		fmt.Scan(&s)
		workArray[i] = s
	}
	for i := 0; i < 3; i ++{
		var a, b int
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for i := 0; i < 10; i ++ {
		fmt.Print(workArray[i], " ")
	}
	fmt.Println("type ok")
}