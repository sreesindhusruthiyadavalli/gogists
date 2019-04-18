package main

import (
	"fmt"
)

var sli []int

func main() {
	//sli := make([]int, 3)
	//for i := 0; i < 3; i++
	while true {
		//_ := range sli {
		//fmt.Println(x)
		var number int
		fmt.Println("Enter an interger:")
		// num, err := fmt.Scanln(&number)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Scanln(&number)
		//sli[i] = number
		sli = append(sli, number)
	}
	fmt.Println(sli)
}
