package main

//go run pr001

import (
	"fmt"
	//ch "pr001/chan"
	//"pr001/p_arr"
	vs "pr001/varStruct"
)

func main() {
	fmt.Println("-------------BEGIN--------------")
	//x := ""
	//fmt.Scan(&x)
	vs.A, vs.B, vs.C = 22, 33, 44 // так можно иниц-ть!
	vs.FFvs1(01, 02, 03)          //a, b, c
	vs.FFvs2()
}
