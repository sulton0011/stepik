// В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

// groupCity := map[int][]string{
// 	10:   []string{...}, // города с населением 10-99 тыс. человек
// 	100:  []string{...}, // города с населением 100-999 тыс. человек
// 	1000: []string{...}, // города с населением 1000 тыс. человек и более
// }

// При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:

// cityPopulation := map[string]int{
// 	город: население города в тысячах человек,
// }

// Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч. 
// была сохранена информация о городах, которые входят в другие группы из groupCity.

// Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation, чтобы в ней была 
// сохранена информация только о городах из группы groupCity[100].

// Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам. По результатам выполнения 
// ничего больше делать не требуется, проверка будет осуществлена автоматически.

package main

import (
	"fmt"
)

func main(){
	groupCity := map[int][]string{
		10:   {`a`,`d`,`e`},
		100:  {`b`,`m`,`t`},
		1000: {`c`,`s`,`y`},
	}
	
	fmt.Println(groupCity)
	
	cityPopulation := map[string]int{
		`a`: 100,//
		`d`: 100,//
		`e`: 1200,//
		
		`b`: 200, 
		`m`: 10,//
		`t`: 300,
			
		`c`: 1200, 
		`s`: 900,//
		`y`: 1300,
	}
//////
	for j, i := range groupCity {
		if j == 10 {
			for _, value := range i{
				if cityPopulation[value] < 10 || cityPopulation[value] >= 100 {
					delete(cityPopulation, value)
				}
			}
		}
		if j == 100 {
			for _, value := range i{
				if cityPopulation[value] < 100 || cityPopulation[value] >= 1000 {
					delete(cityPopulation, value)
				}
			}
		}
		if j == 1000 {
			for _, value := range i{
				if cityPopulation[value] < 1000 {
					delete(cityPopulation, value)
				}
			}
		}
	}
/////////

	fmt.Println(cityPopulation)
	
}