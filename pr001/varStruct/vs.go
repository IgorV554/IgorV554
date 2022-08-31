package vs

import "fmt"

//go run pr001
/* Зарезервированные слова:
break, case, chan, const, continue, default, defer, else, fallthrough, for,
func, go, goto, if, import, interface, map, package, range, return, select,
struct, switch, type, var
byte\uint8:   0...255
int16:  -32768...32767
uint16: 0...65535
int32:  -2 147 483 648...2 147 483 647
uint32: 0...4 294 967 295
int64:  0...18 446 744 073 709 551 615
uint64: –9 223 372 036 854 775 808...9 223 372 036 854 775 807
rune (int32) используется для представления символа Unicode
float32: 1.4*10^-45...3.4*10^38
float64: 4.9*10^-324...1.8*103^08
complex64:  вещественная/мнимая float32  1+2i
complex128: вещественная/мнимая float64

*/
// Глобальные переменные - с большой буквы
var A, B, C byte //

func FFvs1(a byte, b byte, c byte) {
	fmt.Println("Переменные внеш A B C: ", A, B, C) //
	fmt.Println("Переменные внут a b c: ", a, b, c) //
	s := `\n Если текст "в кавычках" то так\n`      // \n Если текст "в кавычках" то так\n
	p := "\n Если текст 'в кавычках' то так\n"      // Если текст 'в кавычках' то так
	fmt.Println(s, p)
}

/* применение в мэйн
vs.A, vs.B, vs.C = 22, 33, 44 // так можно иниц-ть!
vs.FFvs1(01, 02, 03)          //a, b, c
*/

// c маленькой буквы область видимости только этот пакет!
// неявная типизация
var name, age, num, comp = "Tom", 27, 1.05, 4 + 8i

// явная
var (
	name1 string = "D"
	age1  byte   = 0
)

func FFvs2() {

	xType := fmt.Sprintf("%T", name)
	fmt.Println("name (Тип): ", xType)                     // string
	fmt.Println("age (Тип): ", fmt.Sprintf("%T", age))     // int
	fmt.Println("name1 (Тип): ", fmt.Sprintf("%T", name1)) // string
	fmt.Println("age1 (Тип): ", fmt.Sprintf("%T", age1))   // uint8
	fmt.Println("age1 (Тип): ", fmt.Sprintf("%T", num))    // float64
	fmt.Println("age1 (Тип): ", fmt.Sprintf("%T", comp))   // complex128
}
