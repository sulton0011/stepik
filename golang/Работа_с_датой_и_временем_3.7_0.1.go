// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в
// том же формате.

// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем
// вывести на стандартный вывод в том же формате.

// Sample Input: 2020-05-15 08:00:00
// Sample Output: 2020-05-15 08:00:00

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var times string

	times, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	times = strings.ReplaceAll(times, "\n", "")

	times1 , _:= time.Parse("2006-01-02 15:04:05", times)

	if times1.Hour() > 13 {
		times1 = times1.AddDate( 0, 0, 1)
	}

	fmt.Println(times1.Format("2006-01-02 15:04:05"))
}