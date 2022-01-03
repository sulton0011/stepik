// Пришло время для задач, где вы сможете применить полученные знания на практике.

// Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(), которая
// возвращает 3 значения типа пустой интерфейс. Эта функция использует пакеты encoding/json, fmt, и os - не удаляйте
// их из импорта. Скорее всего, вам понадобится пакет "fmt", но вы можете использовать любой другой пакет для записи
// в стандартный вывод не удаляя fmt.

// Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64.
// Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная
// математическая операция). Но такие идеальные случаи будут не всегда, вы можете получить значения других типов,
// а также строка в третьем значении может не относится к одной из указанных математических операций.

// Результат выполнения программы должен быть таков:

// 1. в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической
// операции с точностью до 4 цифры после запятой (fmt.Printf(%.4f));

// 2. если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке
// вида: value=полученное_значение: тип_значения (например: value=true: bool)

// 3. если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим
// операциям, сообщение об ошибке должно иметь вид: неизвестная операция

// Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения
// увидели, что оно содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных
// аргументов значения уже не имеет, а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия
// выполнения задания.

package main

import (
	"fmt"
)

//	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
//	"fmt"           // пакет используется для проверки ответа, не удаляйте его
//	"os"            // пакет используется для проверки ответа, не удаляйте его



func main() {
	var value1, value2, operation interface{} = true, 3.0, 12 // исходные данные получаются с помощью этой функции
    // все полученные значения имеют тип пустого интерфейса



	s1, i := value1.(float64)
	if s2, j :=  value2.(float64); i && j {
		var sum float64
		switch operation{
			case "+": sum = s1 + s2
			case "-": sum = s1 - s2
			case "*": sum = s1 * s2
			case "/": sum = s1 / s2
			default :
			fmt.Println("неизвестная операция")
		}
		fmt.Printf("%.4f\n", sum)
	}else {
		if _, i := value1.(float64); !i {
			fmt.Printf("value=%v: %T", value1, value1)
		}else if _, i := value2.(float64); !i {
			fmt.Printf("value=%v: %T", value2, value2)
		}
	}
}

