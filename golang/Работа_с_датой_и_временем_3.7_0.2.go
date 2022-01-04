// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в 
// 	примере).

// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и 
// большей датами.

// Sample Input: 13.03.2018 14:00:15,12.03.2018 14:00:15
// Sample Output: 24h0m0s

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	s := strings.Split(in.Text(), ",")

	s1, _ := time.Parse("02.01.2006 15:04:05", s[0])
	s2, _ := time.Parse("02.01.2006 15:04:05", s[1])

	if time.Since(s2).Round(time.Second) > time.Since(s1).Round(time.Second){
		fmt.Println(time.Since(s2).Round(time.Second) - time.Since(s1).Round(time.Second))
	}else {
		fmt.Println(time.Since(s1).Round(time.Second) - time.Since(s2).Round(time.Second))
	}
}