// В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной
// жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

// Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
// У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

// Если значение On == false, то оба метода вернут false.

// Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его
// нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

// Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем
// testStruct в функции main, в дальнейшем программа проверит результат.

// Закрывающая фигурная скобка в конце main() вам не видна, но она есть.

// Пакет main объявлять не нужно!

// Удачи!

// Sample Input:
// Sample Output:

package main

import . "fmt"

type S struct {
	On bool
	Ammo int
	Power int
}
func (s *S) Shoot() bool{
	if s.On != true || s.Ammo == 0 {
		return false
	}else {
		s.Ammo -= 1
		return true
	}
}
func (s *S) RideBike() bool{
	if s.On != true || s.Power == 0 {
		return false
	}else {
		s.Power -= 1
		return true
	}
}
func main() {
    testStruct := &S{}
	Println(testStruct)
}