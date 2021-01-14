package main

import (
	"fmt"
)

func evood(d []int) {

	for _, a := range d {
		if a%2 == 0 {
			fmt.Println("the number is even", a)
		} else {
			fmt.Println("the number is odd", a)
		}
	}

}
