package main

//go run pr001

import (
	"fmt"

	/*
			//ch "pr001/chan"
			//"pr001/p_arr"
			//
			//"pr001/types" s "pr001/VarStruct"
		vs "pr001/VarStruct"
		pf "pr001/p_fun"
		"pr001/cycle"
		"pr001/inf"*/
	"pr001/inf"
)

func main() {
	fmt.Println("-------------BEGIN--------------")
	/*x := ""
	fmt.Scan(&x)*/
	//arr := []string{"q", "a", "b", "c", "d", "e"}
	//pf.FF7()
	//
	d := inf.Dat{}
	fmt.Println(d.Str()) // Результат выполнения:
	//
	m := inf.Mat{N1: 10, N2: 10}
	// Это работает и без объявления интерфейса
	fmt.Println(m)       // {10 10}
	fmt.Println(m.Add()) // 20
	fmt.Println(m.Sub()) // 0
	fmt.Println(m.Mul()) // 100
	fmt.Println(m.Div()) // 1
	//
	e := inf.End{}
	fmt.Println(e.Str()) // Конец операции
}
