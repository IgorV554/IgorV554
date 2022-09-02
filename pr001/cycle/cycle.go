package cycle

//go run pr001

import (
	"fmt"
	"log"
)

func FFcy1(x *[]int) {
	for i, v := range *x {
		fmt.Println(i+5, v)
		i++
	}
}

func FFcy2() {
	log.Println("пример Select")
}

/*5 1
6 2
7 3
8 10
9 20
10 30*/

/* применение в мэйне:
arr:=[...]int{1,2,3,10,20,30}
cycle.FFcy1(&arr)
*/
// switch/case с более чем одним условием
func FFcy3() string { //arg ...[]string)  string {
	arg := []string{"q", "a", "b", "c", "d", "e"}
	for _, i := range arg {
		switch i {
		case "a":
			fmt.Println("a: ", i)
		case "b":
			fmt.Println("b: ", i)
		case "c", "d":
			fmt.Println("cd: ", i)
		default:
			{
				fmt.Println("break: ", i)
				break
			}
		}
	}
	return "END"
}

/* в мэйне:
s := cycle.FFcy3()
	fmt.Println(s)
консоль:
-------------BEGIN--------------
break:  q
a:  a
b:  b
cd:  c
cd:  d
break:  e
END
*/
