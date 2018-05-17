package main

import (
	_ "courser-server/go-sql-driver/mysql"
	"database/sql"
)

// DB stores the database connection gloablly.
var DB *sql.DB

// OpenDB opens a database connection
func OpenDB(username string, password string, address string, name string) {
	d, err := sql.Open("mysql", username+":"+password+"@tcp("+address+")/"+name)
	Check(err, true)
	err = d.Ping()
	Check(err, true)
	DB = d
}

// Student is a struct to save data from the Students table.
type Student struct {
	id        string
	firstName string
	lastName  string
	password  string
	classname string	
}

// GetStudents prints a list of all students registered in the database.
func GetStudents() (students []Student) {
	rows, err := DB.Query("SELECT id, first_name, last_name, classname FROM Students")
	Check(err, true)
	defer rows.Close()
	for rows.Next() {
		s := Student{}
		err := rows.Scan(&s.id, &s.firstName, &s.lastName, &s.classname)
		Check(err, true)
		students = append(students, s)
		Log(s.id, s.firstName, s.lastName, s.password)
	}
	return students
}
