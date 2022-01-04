// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам
// дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования
// 	равны 0).

// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате
// UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

// Sample Input: 12 мин. 13 сек.
// Sample Output: Fri May 15 19:28:18 UTC 2020

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)
const now = 1589570165

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	now1 := time.Unix(now, 0).UTC()
	s1 , _ := strconv.Atoi(in.Text()[:2])
	s2 , _ := strconv.Atoi(in.Text()[11:13])

	now1 = now1.Add(time.Minute * time.Duration(s1)).UTC()
	now1 = now1.Add(time.Second * time.Duration(s2)).UTC()
	
	fmt.Println(now1.Format(time.UnixDate))
}