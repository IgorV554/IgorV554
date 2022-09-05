package pf

// массивы срезы указатели на них

import "fmt"

// go run pr001
// -------------------------------------------------------
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
// -------------------------------------------------------
// передача массива по указателю
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
// -------------------------------------------------------
// аргумент
func FF3(a int, s string) (int, string) {
	return a, s
}

/* в мэйне
var (
	a int = 0
	c string
)
a++
c += "0x0F"
fmt.Println(pf.FF3(a, c)) // 1 0x0F
*/
// -------------------------------------------------------
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

// -------------------------------------------------------
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

// -------------------------------------------------------
// функция внутри функции - вот именно эта - анонимная ф-я
func FF6() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1)) // 2
}

// -------------------------------------------------------
// Функция как параметр другой ф-и
func Add(x, y int) int {
	return x + y
}

func Mul(x, y int) int {
	return x * y
}

func Sub(x, y int) int {
	return x - y
}

// -------------------------------------------------------
// Функция объявлена как переменная - f
func FF7() {
	var f func(int, int) int = Add
	fmt.Println(f(7, 8)) // 15
	f = Mul
	fmt.Println(f(5, 6)) // 30
}

// -------------------------------------------------------
// Функция как аргумент
func FF8(x int, y int, function func(int, int) int) int {
	return function(x, y)
}

/* вызов в мэйне:
pf.FF7()
ret := pf.FF8(3, 8, pf.Add) // 11
fmt.Println(ret)
ret = pf.FF8(3, 8, pf.Mul) // 24
fmt.Println(ret)
*/

// Пример - сумма всех нечетных чисел среза
func Even(x int) bool {
	return x%2 == 0
}
func FcriterySum(arr []int, critery func(int) bool) int {
	ret := 0
	for _, v := range arr {
		if critery(v) { // если критерий вернул True
			ret += v
		}
	}
	return ret
}

/* в мэйне:
Ar := []int{1, 2, 3, 4, 5, 6, 7}
Sum := pf.FcriterySum(Ar, pf.Even)
fmt.Println("Сумма четных: ", Sum) // 12
*/
// -------------------------------------------------------
// Функция возвращающая функцию
func ReturnFunc(n int) func(int, int) int {
	switch n {
	case 0:
		fmt.Printf("Add: ")
		return Add
	case 1:
		fmt.Printf("Mul: ")
		return Mul
	default:
		fmt.Printf("Sub: ")
		return Sub
	}
}

/* в мэйне:
----- var f func(int,int) = pf.ReturnFunc(0) - вот это писать не надо! сразу так:
	f := pf.ReturnFunc(0)
	fmt.Println(f(10, 1)) // Add: 11
	f = pf.ReturnFunc(1)
	fmt.Println(f(10, 1)) // Mul: 10
	f = pf.ReturnFunc(2)
	fmt.Println(f(10, 1)) // Sub: 9
*/

// -------------------------------------------------------
// Анонимная функция - ее можно объ внутри ф-и
/* 	Преимуществом анонимных функций является то,
что они имеют доступ к окружению,
в котором они вызываются
*/
// а тут я возвращаю две функции!
func FF9() (func(int, int) int, func(int, int) int) {
	f := func(x, y int) int { return x + y }
	// можно и так:
	var g func(x, y int) int = func(x, y int) int { return x - y }
	//fmt.Println(f(6, 7))
	//fmt.Println(g(6, 7))
	return f, g
}

/* в мэйне:
a, b := pf.FF9() // - ЭТО РАБОТАЕТ!
fmt.Println(a(10, 4)) // 14
fmt.Println(b(10, 5)) // 5
*/

// -------------------------------------------------------
// Анонимная функция как аргумент функции
func Afun_as_arg(x int, y int, operator func(int, int) int) int { // анонимная поступает на вход в мэйне
	return operator(x, y)
}

/* в мэйне:
r := pf.Afun_as_arg(20, 5, func(a int, b int) int { return a + b })
fmt.Println(r) // 25
r = pf.Afun_as_arg(20, 5, func(a int, b int) int { return a - b })
fmt.Println(r) // 15
r = pf.Afun_as_arg(20, 5, func(a int, b int) int { return a * b })
fmt.Println(r) // 100
*/
// -------------------------------------------------------
// Анонимная функция как результат функции - это кстати было выше
func FAfun(n int) func(int, int) int {
	switch n {
	case 0:
		return func(x int, y int) int {
			fmt.Printf("Anon_add: ")
			return x + y
		}
	case 1:
		return func(x int, y int) int {
			fmt.Printf("Anon_sub: ")
			return x - y
		}
	default:
		return func(x int, y int) int {
			fmt.Printf("Anon_mul: ")
			return x * y
		}
	}
} /* в мэйне:
r := pf.FAfun(0)     // add
fmt.Println(r(4, 5)) //	Anon_add: 9
r = pf.FAfun(1)      // sub
fmt.Println(r(4, 5)) // Anon_sub: -1
r = pf.FAfun(2)      // mul
fmt.Println(r(4, 5)) // Anon_mul: 20
*/
// -------------------------------------------------------
/* замыкание - для него используют анонимные функции, потому что они
имеют доступ к своему окружению */
func Fsquare(x int) func() int {
	return func() int {
		x++
		return x * x
	}
}

/* Рекурсия!
f := pf.Fsquare(2) // рекурсия
fmt.Println(f())   // 9
fmt.Println(f())   // 16
fmt.Println(f())   // 25
*/
