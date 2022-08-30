package main

//go run pr001

import (
	"fmt"
	ch "pr001/chan"
)

func main() {
	fmt.Println("-------------BEGIN--------------")
	//x := ""
	//fmt.Scan(&x)
	//pointer.FF1(&x)
	/*
		// перевод дес в hex и печать в консоль
		var arr = []int{1, 2, 4, 8, 16, 32, 64, 128, 255}
		for i := 0; i < len(arr); i++ {
			hex.FF2(arr[i]) // 0x1 0x2 0x4 0x8 0x10 0x20 0x40 0x80 0xFF
		}*/

	//	pointer.FF2()
	// !передача данных в канал без запущеной го-рутины не работает!
	sa := make(chan string)
	go ch.Routine1(sa)
	sa <- "ЭТО КАНАЛ"
	//go Routine1(sa)
	//FFch1(sa)

}
