package types

//go run pr001
/*
type mile(новый тип) int

type имя_структуры struct{
    поля_структуры
}

*/
import "fmt"

type Size int32

func FFt1() {
	fmt.Printf("----------ТИПЫ-----------")
	var W, L, H Size = 10, 20, 30
	{
		W += +10 // тоже что W += 10 - 20
		L += -10 // 10
		H /= 2   // 15
	} // так можно!
	fmt.Println(W, " ", L, " ", H, " ")
}

type Notes string // исп вне пакета - с большой буквы!

func PrintNotes1(N *[]Notes) { // тут понятно что тип Notes это массив
	for _, v := range *N {
		fmt.Println(v)
	}
	/*
	   Note1
	   Note2
	   Note3
	   Note4
	*/
}

/* применение в мэйне - примечание 1:
// из-за того что мы назвали тип Notes (хотя он string) это не заработает
	NL:= []string{"Note1","Note2","Note3","Note4"}
	types.PrintNotes1(&NL)
	// надо так (не string а именно Notes!):
	NL1 := []types.Notes{"Note1", "Note2", "Note3", "Note4"}
	types.PrintNotes1(&NL1)
*/

type Notez []string

func PrintNotes2(N *Notez) { // тип Notez инкапсулирует в себе информацию что это массив
	for _, v := range *N {
		fmt.Println(v)
	}
}

/*
Note-1
Note-2
Note-3
Note-4
*/
/* применение в мэйне - примечание 2:
// тут понятно что NL1 это массив
NL1 := []types.Notes{"Note1", "Note2", "Note3", "Note4"}
types.PrintNotes1(&NL1)

// при создании NL2 в типе инкапсулировано что это массив
var NL2 types.Notez = types.Notez{"Note-1", "Note-2", "Note-3", "Note-4"}
types.PrintNotes2(&NL2)
*/
