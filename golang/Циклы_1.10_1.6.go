// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

// Входные данные
// Программа получает на вход два числа. Гарантируется, что цифры в числах не
// повторяются. Числа в пределах от 0 до 10000.

// Выходные данные
// Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
// Цифры выводятся в порядке их нахождения в первом числе.
// Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше,
// затем вернетесь к этой задаче позже.

// Sample Input:
// 564 8954

// Sample Output:
// 5 4

package main

import "fmt"

func main() {
 var num1, num2 int

 fmt.Scan(&num1, &num2)
 n1 := split(num1)
 n2 := split(num2)

 for _, i := range n1 {
  for _, j := range n2 {
   if i == j {
    fmt.Print(i, " ")
   }
  }
 }

}

func split(num int) []int {
 var t []int
 for num > 0 {
  t = append(t, num%10)
  num = num / 10
 }
 var nums []int
 for i := len(t) - 1; i >= 0; i-- {
  nums = append(nums, t[i])
 }
 return nums
}