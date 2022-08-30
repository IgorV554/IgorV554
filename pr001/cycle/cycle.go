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
