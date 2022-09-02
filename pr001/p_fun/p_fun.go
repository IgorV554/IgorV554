package pf

// массивы срезы указатели на них

import "fmt"

//go run pr001

// передача среза по указателю
func FF1(x *[]int) {
	fmt.Println("указатель на срез", x) // указатель на срез &[1 2 3]
	// почему-то это не изменило элементы среза!!
	for _, i := range *x {
		i = i + 1
	}
	fmt.Println("изменен", x) // изменен &[1 2 3] <не был!!>
}

/* в мэйне
ar := []int{1, 2, 3}
pr := make([]int, 3)
pf.FF1(&ar) // изменен &[1 2 3]
pf.FF1(&pr) // изменен &[0 0 0]
// mr :=[...]int{1,2,3}pf.FF1(&mr) // так нельзя - в го нет перегрузки
*/

func FF2(x *[5]int) { // (x *[...]int) нельзя!
	fmt.Println("указатель на массив", x) // указатель на массив &[1 2 3 4 5]

	// почему-то это не изменило элементы массива!!
	for _, i := range *x {
		i = i + 1
	}
	fmt.Println("не изменен", *x) // изменен &[1 2 3 4 5] <не был!!>

	// а тут он стал изменен:
	for i := 0; i < len(x); i++ {
		x[i] += 1
	}
	fmt.Println("изменен", *x) // изменен [2 3 4 5 6]!!
}

/*
	a := [5]int{1, 2, 3, 4, 5}
	pf.FF2(&a)
*/

// несколько аргументов
func FF3(a int) int {
	return a
}

/* в мэйне
var (
	a int = 0
	c string
)
a++
c += "0"
fmt.Println(pf.FF3(10)) // 10
*/

// возврат двух параметров
func FF4(a int, c string) (int, string) {
	return (a + 1), (c + "__")
}

/* в мэйне
var (
	a int    = 10
	c string = "VVV"
)
fmt.Println(pf.FF4(10, "S")) // 11 S__
a, c = pf.FF4(a, c)
fmt.Println(a, c) // 11 VVV__
*/

// неопределенное число аргументов
func FF5(Arg ...int) int {
	fmt.Println(Arg)
	var Sum int = 0
	for _, v := range Arg {
		Sum = Sum + v
	}
	return Sum
}

/*
	var a, b, c, d int = 20, 21, 22, 23
	pf.FF5(a, b)    // [20 21]
	pf.FF5(b, c, d) // [21 22 23]
	fmt.Println(pf.FF5(a, b, c, d)) //[20 21 22 23]  86
*/

// функция внутри функции
func FF6() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1)) // 2
}

// анонимная функция

// замыкание - для него используют анонимные функции
