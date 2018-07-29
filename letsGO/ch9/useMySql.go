package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------")
	defer db.Close()
	var result sql.Result
	//向数据库中插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy) values(?,?,?)", "张三", 19, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("+++++++++++++++")
	lastId, _ := result.LastInsertId()
	fmt.Println("新插入的数据 ID 为", lastId)
	var row *sql.Row
	//返回一行数据
	row = db.QueryRow("select * from person")
	var name string
	var id, age int
	var isBoy bool
	//取数据进行显示
	err = row.Scan(&id, &name, &age, &isBoy)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id, "\t", name, "\t", age, "\t", isBoy) //再插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy)values(?,?,?)", "王红", 18, false)
	fmt.Println("=====================")
	var rows *sql.Rows
	rows, err = db.Query("select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name string
		var id, age int
		var isBoy bool
		rows.Scan(&id, &name, &age, &isBoy)
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	}
	rows.Close()
	// 最后，清空表 db.Exec("truncate table person")
}
