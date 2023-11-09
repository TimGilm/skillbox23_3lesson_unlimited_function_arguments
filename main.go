// модуль 22 урок 3 неограниченное количество аргументов внутри функции
package main

import "fmt"

func main() {
	//1fmt.Println(generalSum(2, 3, 4))
	//2fmt.Println(conditionalCalculate("*", 3, 4))
	anyObjects(1, "222", [2]int{1, 2})
}

//1напишем функцию принимающую любое количество аргументов
/*func generalSum(s ...int) int { //такая запись представляет s  как массив неопределенной длины (срез)
	//т.к. s срез мы можем обращаться к нему как к срезу
	fmt.Println(s[0])
	fmt.Println(s[1])
	ss := 0
	for _, v := range s {
		ss += v
	}
	return ss
}*/

//2напишем функцию, принимающую строку и любое количество аргументов
/*func conditionalCalculate(op string, x ...int) int {
	var result int
	if op == "+" {
		for _, v := range x {
			result += v
		}
	} else if op == "*" {
		result = 1
		for _, v := range x {
			result *= v
		}
	}
	return result
}*/

// напишем функцию принимающую на вход и строки и массивы и др. любые агрументы
func anyObjects(a ...interface{}) {
	fmt.Printf("%+v\n", a)
}
