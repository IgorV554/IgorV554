package tm

import (
	"fmt"
	"time"
)

//go run pr001
/* Работа с таймером */

// ---------------------------------------------------------------------
// Формат вывода времени
func Frt01() {
	fmt.Println("----------RT-----------")
	// по даташиту на язык https://pkg.go.dev/time#pkg-constants
	nt := time.Now().Format(time.RFC822)
	fmt.Println(nt) // 07 Sep 22 10:23 +0300
	//nt = time.Now() // ошибка!
	fmt.Println(time.Now()) // 2022-09-07 10:23:41.0883799 +0300 +0300 m=+0.300463701
	nt = time.Now().Format(time.RFC850)
	fmt.Println(nt) // Wednesday, 07-Sep-22 10:25:27 +0300
	nt = time.Now().Format(time.ANSIC)
	fmt.Println(nt)                               // Wed Sep  7 10:26:52 2022
	nt = time.Now().Format("02.01.2006 15:04:05") // по нашему
	fmt.Println(nt)                               // 07.09.2022 11:53:29
}

/* применение в мэйне: rt.Frt01() */

// ---------------------------------------------------------------------
// Задержка
func Frt02() {
	fmt.Println("----------RT-----------")
	nt := time.Now() // создаем переменную для вывода форматированного времени
	tf := time.Now().Second()
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		nt = time.Now().Local().UTC()
		fmt.Println(i, nt) // 9 2022-09-07 07:36:03.693596 +0000 UTC
		tf = time.Now().Second()
		fmt.Println(i, tf) // 50 - выводятся секудны
	}
}

/* применение в мэйне: rt.Frt02() */

// ---------------------------------------------------------------------
// Тикер №1
func Ftiker_01() {
	// создали тикер
	tiker := time.NewTicker(500 * time.Millisecond)

	// создаем горутину
	go func() {
		fmt.Println("----go----")
		// range автоматически останавливает цикл, когда канал будет закрыт
		for t := range tiker.C { // создаем канал t
			fmt.Println("секунды: ", t.Second()) // секунды:  51 (не от 0 а текущее время!)
		}
	}()

	// пауза на 5 секунд
	time.Sleep(5 * time.Second)
	tiker.Stop() // останавливаем созданный тикер
}

/* в мэйне:
rt.Ftiker_01()
----go----
секунды:  15
секунды:  16
секунды:  17
секунды:  18
секунды:  19
*/

// ---------------------------------------------------------------------
// Тикер №2
func Ftiker_02() {
	fmt.Println("----------RT-----------")
	/*Ticker содержит канал, который поставляет "тики" с интервалами.
	Создаем тикер: раз в 0.5 сек */
	ticker := time.NewTicker(500 * time.Millisecond)
	// задаем выполнение останова в любом случае (видимо это когда будет вызван деструктор)
	defer ticker.Stop()
	// создаем  канал тикера лог типа
	done := make(chan bool)
	// запуск горутины - по завершению ее работы в канал запишется true
	go func() {
		time.Sleep(3 * time.Second) // пауза за 3 сек
		done <- true
	}()
	for {
		select {
		case <-done: // если стало true - это через 3 сек - конец
			fmt.Println("Готово!")
			return
		case t := <-ticker.C:
			fmt.Println("Текущее время: ", t)
		}
	}
}

/* применение в мэйне:  */
