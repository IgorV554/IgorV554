package inf

// интерфейсы

import "fmt"

//go run pr001

/*
Интерфейс это абстракция, декларирование обобщенной
реализации функционала без привязки к конкр типу.
Интерфейс реализуется неявно
*/

// Первая часть. Привязка метода к структуре
type Person struct {
	Name string // имя
	Numb int    // телефон
}

func (p *Person) tName() { // не функция а метод стурктуры
	fmt.Println("Имя: ", p.Name)
}
func (p *Person) tNumb() { // метод
	fmt.Println("Телефон: ", p.Numb)
}

func Finf01() { // указатель на структуру - меняем структуру а не ее экземпляр
	fmt.Println("--------- Интерфейсы 1 ч ---------") // 10000
	p := Person{"Вася", 999}
	fmt.Println("", p.Name, p.Numb)
	p.tName() // Имя:  Вася
	p.tNumb() // Телефон:  999
}

/* применение в мэйн  inf.Finf01()  */
// ------------------------------------------------------------------
// Пример того как два тазных типа переменных в структурах удовлетворяют одному интерфейсу
// объявляем интерфейс
type Stream interface {
	Add() int32
	Sub() int32
	Mul() int32
	Div() int32
	Str() string
}

// и структуру №1
// Структура удовлетворяет интерфейсу потому что у нее есть тип int32
type Mat struct {
	N1, N2 int32
}

// связываем структуру и интерфейс методом на основе структуры
// func (s *Mat) Str() string { return "Результат: " } // компилятор ругается что в структуре Mat нет типа string
func (m *Mat) Add() int32 { return m.N1 + m.N2 }
func (m *Mat) Sub() int32 { return m.N1 - m.N2 }
func (m *Mat) Mul() int32 { return m.N1 * m.N2 }
func (m *Mat) Div() int32 { return m.N1 / m.N2 }

/* в мэйне:
m := inf.Mat{N1: 10, N2: 10}
fmt.Println(m) // {10 10}
fmt.Println(m.Add()) // 20
fmt.Println(m.Sub()) // 0
fmt.Println(m.Mul()) // 100
fmt.Println(m.Div()) // 1
// Примечание: если закомментировать объявл. интерфейса то будет тоже самое
*/
// Структура №2 тоже удовлетворяет интерфейсу - и там и там есть тип string
type Dat struct {
	S string
}

type End struct {
	S string
}

// два метода разных структур:
func (s *Dat) Str() string { return "Результат выполнения: " }
func (s *End) Str() string { return "Конец операции" }

/* в мэйне:
	d := inf.Dat{}
	fmt.Println(d.Str()) // Результат выполнения:

	e := inf.End{}
	fmt.Println(e.Str()) // Конец операции
}
*/
