package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test?charset=urf8")
	if err != nil {
		panic
	}
	rows, err := db.Query("SELECT name FROM user WHERE age=?", age)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
