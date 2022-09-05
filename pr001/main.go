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
		"pr001/cycle"*/
	pf "pr001/p_fun"
)

func main() {
	fmt.Println("-------------BEGIN--------------")
	/*x := ""
	fmt.Scan(&x)*/
	//arr := []string{"q", "a", "b", "c", "d", "e"}
	//pf.FF7()

	f := pf.Fsquare(2) // рекурсия
	fmt.Println(f())   // 9
	fmt.Println(f())   // 16
	fmt.Println(f())   // 25
}
