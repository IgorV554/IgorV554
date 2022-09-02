package p_arr

// массивы срезы указатели на них

import "fmt"

//go run pr001

/*
массивы:
var numbers [5]int
var numbers [5]int = [5]int{1,2,3,4,5}
numbers := [5]int{1,2,3,4,5}
numbers2 := [...]int{1,2,3}

срезы:
numbers := []int{1,2,3,4,5}
numbers := []int{1,2,3}
var numbers <свой тип> = <свой тип>{1,2,3} - на основе своего типа данных
*/

func FFarr1(mArr *[]int) {
	// указатель и оператор взятия адреса - переменная
	// печатаем каждый элемент среза
	fmt.Println("указатели и срезы - 1") //

	for _, value := range *mArr {
		fmt.Printf("%v\n", value)
	}
	/*указатели и срезы - 1
	  1
	  2
	  3
	  10
	  20
	  30*/
	fmt.Println("указатели и срезы - 2") //
	// добавляем в срез новые
	/*	for i := 4; i < 10; {
		*mArr = append(*mArr, i)  - ПК жестко завис!!!
	} */
	*mArr = append(*mArr, 44)
	fmt.Println("Новый срез:", *mArr) // Новый срез: [1 2 3 10 20 30 44]
	/* не компилится! fmt.Println("", mArr[2:5]) // оператор среза */
}

/* применение в мэйн
arr := []int{1, 2, 3, 10, 20, 30}
p_arr.FFarr1(&arr) // массив нельзя передать в ф-ю ожидающую срез!!
*/

func FFarr2(m []int) {
	fmt.Println("Пример 2: создаем срез") //
	s := m[:]                             // таким образом можно создать срез, содержащий весь массив
	// s[]:  [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("s[]: ", s)                    //
	fmt.Println("срез среза s[2:5]: ", s[2:5]) // срез среза s[2:5]:  [3 4 5]
	fmt.Println("срез среза s[:5]: ", s[:5])   // срез среза s[:5]:  [1 2 3 4 5]
	fmt.Println("срез среза s[5:]: ", s[5:])   // срез среза s[5:]:  [6 7 8 9 10]
	fmt.Println("срез среза s[:]: ", s[:])     // срез среза s[:]:  [1 2 3 4 5 6 7 8 9 10]
}

/* применение
m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // создаём
p_arr.FFarr2(m)
*/

// изучаем make что такое емкость
func FFarr3() {
	a := make([]int, 2, 3)             // длина  емкость
	fmt.Println(a, len(a), cap(a))     //
	a = append(a, 1)                   // добавим один
	fmt.Println("Добавили 1эл", a)     // [0 0 1]
	a = append(a, 1)                   // добавим один
	fmt.Println("Добавили еще 1эл", a) // [0 0 1 1] - ?

}
