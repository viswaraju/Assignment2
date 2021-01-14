package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* type Table struct {
	ID, NAME, AGE, COUNTRY string
}

type Org struct {
	OrgTable []Table
}
*/
/* const (
	x = 6
	y
	z
) */
/* type dimensions interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

type square struct {
	side float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}
*/

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenverifyMiddleWare(protectedEndpoint)).Methods("GET")
	log.Println("Listening to :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

	//assignments.Day8to12Assignment2()

	/* abc := map[string]int{
		"One": 1,
		"two": 2,
	}

	for a, b := range abc {
		fmt.Println(a, b)

	}
	*/
	/* args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [uesrname] [password]")
		return

	}

	u, p := args[1], args[2]

	if u == "jack" && p != "1888" {
		fmt.Println("Access granted for", u, p)
	} else if u == "inanc" && p == "1879" {
		fmt.Println("Access granted for", u, p)
	} else {
		fmt.Println("Access denied", u, p)
	}


	*/
	//personal.Stru()
	/* colours := make(map[string]string)
	//var colours map[string]string
	colours["white"] = "kkkkkk"
	delete(colours, "white")

	fmt.Println(colours) */
	/* cards := assignment.Newdeck()

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
	cards.Print(

	func(b aas) num(){

	}




	/* cards := assignment.Newdeck()
	cards.Print() */
	/* d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, a := range d {
		fmt.Println(a)
	} */
	/* d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evood(d) */
	//personal.Day8to12Assignment2()
	/* personal.GoSQLProgram()
	personal.Checkstatus() */

	/* data, err := ioutil.ReadFile("assignment4\\file1.txt")

	if err != nil {
		fmt.Println("couldnt read the file")
	} else {
		a := string(data)
		b := []string{}
		b = append(b, a)
		fmt.Println(b)

		sum := b[0] + b[1]
		fmt.Println(sum) */
	//}
	//assignment3.Interface()
	//assignment3.Json() */
	//personal.Checkstatus()

	//fmt.print(b)
	//practise.Assignment1()
	/* var s []int

	for {
		var num int
		fmt.Println("Enter any number:")
		fmt.Scanln(&num)
		//fmt.Println(s)
		//var a string
		if num == 0 || num == 1 {
			break
		} else {
			s = append(s, num)
		}

	} */
	//s = s[:len(s)-1]
	//fmt.Println(s)
	//assignment.Assignment1()
	/* for {
	fmt.Println("Enter any number:")
	var s []string
	var num string
	//var q string
	s = append(s, num)
	fmt.Scanln(&num)
	s = append(s, num)
	fmt.Println(s) */

	/* if num == "Q" {
		break

	} */

	//}
	/* s = append(s, num)
	fmt.Println(s) */

	//fmt.Println("Welcome to Go training")

	//"assignment 1:"
	/* num := 14.65478
	fmt.Printf("%.2f", num) */

	//fmt.Println("-----------------------------------------------")

	//"assignment 2:"

	//"Case : 1"

	/* fmt.Println("Input 1 - Enter any string:")
	var num string
	fmt.Scanln(&num)
	fmt.Printf("%T", num)
	if s, err := strconv.ParseInt(num, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} */

	//fmt.Println("-----------------------------------------------")

	//"assignment 3:"

	/* text := "go-lang-training"
	text = strings.ReplaceAll(text, "-", " ")
	fmt.Println(strings.Title(text)) */

	fmt.Println("-----------------------------------------------")

	//"assignment 4:"
	/* fmt.Print(x, y, z) */

	fmt.Println("-----------------------------------------------")

	//"assignment 5:"
	/* val := 20
	newval := 30
	var ptr *int
	ptr = &val
	fmt.Println("previous value of val", *ptr)
	*ptr = newval
	fmt.Println("New value of val", *ptr)
	*/

	fmt.Println("-----------------------------------------------")

	//"assignment 6:"
	/* s := "1Ax3 7y Bk"
	fmt.Println("%T", s)
	//fmt.Println(utf8.RuneCountInString(s))
	ss := []rune(s)
	fmt.Println(ss)

	//fmt.Println(unicode.IsLetter(ss[2]))
	a := 0
	b := 0
	c := 0
	d := 0
	for i := 0; i < len(ss); i++ {
		if unicode.IsDigit(ss[i]) == true {
			a++
		} else if unicode.IsSpace(ss[i]) == true {
			b++
		} else if unicode.IsLower(ss[i]) == true {
			c++
		} else if unicode.IsUpper(ss[i]) == true {
			d++
		} else {
			return
		}

	}
	fmt.Println("The number of digits is:", a)
	fmt.Println("The number of spaces is:", b)
	fmt.Println("The number of Lower case letters are:", c)
	fmt.Println("the number of Uppercase letters are:", d) */

}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup invoked")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked")
}

func TokenverifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenverifyMiddleWare invoked")
	return nil
}
