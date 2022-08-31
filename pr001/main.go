package main

//go run pr001

import (
	"fmt"
	/*
	   //ch "pr001/chan"
	   //"pr001/p_arr"*/
	//vs "pr001/varStruct"
	"pr001/types"
)

func main() {
	fmt.Println("-------------BEGIN--------------")
	//x := ""
	//fmt.Scan(&x)

	// тут понятно что NL1 это массив
	NL1 := []types.Notes{"Note1", "Note2", "Note3", "Note4"}
	types.PrintNotes1(&NL1)

	// при создании NL2 в типе инкапсулировано что это массив
	var NL2 types.Notez = types.Notez{"Note-1", "Note-2", "Note-3", "Note-4"}
	types.PrintNotes2(&NL2)
}
