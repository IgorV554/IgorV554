package ch

//go run pr001

import "fmt"

// рутина для запуска в мэйне
func Routine1(c chan string) {
	fmt.Printf("Данные канала: %v", <-c)
}

/* применение в мэйне:
// !передача данных в канал без запущеной го-рутины не работает!
sa := make(chan string)
go ch.Routine1(sa)
sa <- "ЭТО КАНАЛ" // напечтало: Данные канала: ЭТО КАНАЛ
*/
