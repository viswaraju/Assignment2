package practise

import (
	"fmt"
	"sort"
)

func Assignment() {
	var s []string

	for {

		//fmt.Println(s)
		var num string
		fmt.Println("Enter any string:")
		fmt.Scanln(&num)
		s = append(s, num)
		if num == "q" || num == "Q" {
			break
		}

	}
	s = s[:len(s)-1]
	fmt.Println(s)
}

func Assignment2() {
	var rows int
	var sum int
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)
	for i := 0; i < rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print("  ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				sum = 1
			} else {
				sum = sum * (i - k + 1) / k
			}
			fmt.Printf("%d ", sum)
		}
		fmt.Println(" ")
	}
}

func Assignment3() {

	type Table struct {
		Id      string
		Name    string
		Age     string
		Country string
	}

	type Org struct {
		OrgTable []Table
	}
	e := Org{
		OrgTable: []Table{
			Table{
				Id:      "ID",
				Name:    "Name",
				Age:     "Age",
				Country: "Country",
			},
			Table{
				Id:      "1100",
				Name:    "Jhon",
				Age:     "25",
				Country: "USA",
			},
			Table{
				Id:      "1101",
				Name:    "Ralph",
				Age:     "30",
				Country: "USA",
			},
			Table{
				Id:      "1102",
				Name:    "Rama",
				Age:     "32",
				Country: "INDIA",
			},
		},
	}
	fmt.Println(e.OrgTable[0])
	fmt.Println(e.OrgTable[1])
	fmt.Println(e.OrgTable[2])
	fmt.Println(e.OrgTable[3])

}

func Assignment4() {
	fmt.Println("Sorting Integers")
	var integers = []int{10, 50, 12, 18, 70, 11}
	fmt.Println(integers)
	fmt.Println("Are they sorted:", sort.IntsAreSorted(integers))
	sort.Ints(integers)
	fmt.Println("Intergers are sorted:", integers)
	fmt.Println("------------------")
	fmt.Println("Sorting Float")
	var float = []float64{17.2, 17.1, 23.2, 10.6, 22.6}
	fmt.Println(float)
	fmt.Println("Are they sorted:", sort.Float64sAreSorted(float))
	sort.Float64s(float)
	fmt.Println("Float are sorted:", float)
	fmt.Println("------------------")
	fmt.Println("Sorting Strings")
	var strings = []string{"GOLANG", "Python", "C++", "Java", ".net", "HTML"}
	fmt.Println(strings)
	fmt.Println("Are they sorted:", sort.StringsAreSorted(strings))
	sort.Strings(strings)
	fmt.Println("Intergers are sorted:", strings)
}

func Assignment5() {
	progLaunguages := []string{"Java", "Golang", "Python", "c++", "HTML"}
	fmt.Println(progLaunguages)
	webFrameworks := []string{"Angular", "MeteorJS", "ExpressJs"}
	fmt.Println(webFrameworks)
	progLaunguages = append(progLaunguages, webFrameworks...)
	fmt.Println(progLaunguages)
}
