package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Do not remove. The driver is used during runtime
)

var (
	server   = "localhost"
	port     = "5432"
	user     = "postgres"
	database = "postgres"
)

// Employee ...
type Employee struct {
	Name        string
	ID          int
	Role        string
	JoiningDate time.Time
}

type Address struct {
	ID          int
	Add_name    string
	Street_name string
	City        string
	State       string
	Country     string
	Pincode     string
}

// GetEmpData ...
func GetEmpData(db *sql.DB) (empData []Employee, err error) {

	inputquery := "select emp_id, emp_name, emp_role, joining_date from public.employee "

	rows, err := db.Query(inputquery)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	count := 0

	// cols, err := rows.Columns()
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Role, &emp.JoiningDate)
		if err != nil {
			fmt.Println("Error reading rows" + err.Error())
			return nil, err
		}

		// fmt.Printf("%v, \n", rows)
		empData = append(empData, emp)
		count++
	}
	return empData, nil
}

func Adddata(db *sql.DB) (addData []Address, err error) {

	inputquery :=

		"select emp_id,address_id,street_name, city, state, country, pincode from public.address"

	rows, err := db.Query(inputquery)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	fmt.Print(cols)
	count := 0
	for rows.Next() {
		var addr Address
		err := rows.Scan(&addr.ID, &addr.Add_name, &addr.Street_name, &addr.City, &addr.State, &addr.Country, &addr.Pincode)

		if err != nil {
			fmt.Println("Error reading rows 2: " + err.Error())
			return nil, err
		}
		//fmt.Printf("ID: %d, Name: %s, role: %s, \n", emp_id, emp_name, emp_role)

		//fmt.Printf("%v, \n", rows)

		//fmt.Printf("%v, \n", rows)
		addData = append(addData, addr)
		count++
	}

	return addData, nil
}

// GoSQLProgram ...
func GoSQLProgram() {
	fmt.Println("*** GoSQLProgram ***")

	conn, err := Connection()
	if err != nil {
		fmt.Println("Open connection failed:", err.Error())
		return
	}
	defer conn.Close()
	// Read employees
	emp, err := GetEmpData(conn)
	if err != nil {
		fmt.Println("ReadEmployees failed:", err.Error())
		return
	}

	for _, a := range emp {
		fmt.Println(a)
	}
	//fmt.Printf("Read %d rows successfully.\n", len(emp))
	// Read Addrress
	add, err := Adddata(conn)
	if err != nil {
		fmt.Println("ReadEmployees failed:", err.Error())
		return
	}

	for _, b := range add {
		fmt.Println(b)
	}

}

// Connection ...
func Connection() (con *sql.DB, err error) {

	//dat, _ := ioutil.ReadFile("files\\test.txt")
	password := "qwerty"
	// fmt.Println(password)

	dsn := ``

	dsn += "sslmode=disable "
	dsn += "host=" + server + " "
	dsn += "Application_Name=" + "SQLProgram" + " "
	dsn += "database=" + database + " "
	dsn += "user=" + user + " "
	dsn += "port=" + port + " "
	dsn += "password=" + password

	// fmt.Println(dsn)
	con, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return con, nil
}
