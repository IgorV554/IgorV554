package pointer

import "fmt"

//go run pr001

func FF1(x *string) {
	// указатель и оператор взятия адреса - переменная
	pa := x
	fmt.Println("Адрес ячейки: ", pa) // 0xc00008a050
	fmt.Println("Значение: ", *pa)    // что ввели
	// узнаем какой тип у переменной pa
	fmt.Println("Тип переменной указателя: ", fmt.Sprintf("%T", pa)) // тип *string
}

/* применение
x := ""
fmt.Scan(&x)
pointer.FF1(&x)
*/
