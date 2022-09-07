package rt

import (
	"fmt"
	"time"
)

//go run pr001

/* Работа с горутинами */

// ----------------------------------------------------------
// канал, открытие и ф-я defer
func Frt01() {
	defer fmt.Println("Выполнение defer")
	//
	fmt.Printf("----------ROUTINE-----------")
	fmt.Println(time.Second) // time.Now().Second()
	var ch chan int
	// обработка ошибки закрытия нулевого канала
	close(ch)
}

/*
	применение в мэйне:

rt.Frt01()
// Выход:
---------ROUTINE-----------1s
Выполнение defer
PANIC была вызвана!
*/
// Последовательность работы defer
func Frt02() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}

/* Выход:
three
two
one
*/
//-------------------------------------------------------------
//  Ошибка закрытия нулевого канала и обработка ошибки
func Frt03() {
	var ch chan int
	//close(ch) // panic! так не правильно
	fmt.Println(res(ch)) // ERR
}

func Frt04() {
	ch := make(chan int, 1) // вот так ошибки закрытия не будет
	fmt.Println(res(ch))    // OK, closed
	/* Выход: ERR2:  0x9083e0 */
}

// Проверка канала на то что он не пустой
func res(c chan int) (r string) {
	if c != nil {
		close(c) // так правильно
		r = "OK, closed"
	} else {
		r = "ERR"
	}
	return r
}
