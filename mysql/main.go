package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type movie struct {
	id          int64
	name        string
	publishTime time.Time
}

var db *sql.DB
var recordID int64

func init() {
	var err error
	db, err = sql.Open("mysql", "root:chinaren@tcp(127.0.0.1:3306)/mygwp?charset=utf8&parseTime=true&loc=Local")
	checkErr(err)
}

func main() {
	insert()
	update()
	query()
	delete()
}

// 插入数据
func insert() {
	stmt, err := db.Prepare("insert into movie(name, publish_time) values(?, ?)")
	checkErr(err)
	defer stmt.Close()

	//res, err := stmt.Exec("name", "2016-03-06")
	res, err := stmt.Exec("name", time.Now())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	recordID = id
}

// 更新数据
func update() {
	stmt, err := db.Prepare("update movie set name=? where id=?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec("电影", recordID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

// 查询数据
func query() {
	rows, err := db.Query("select * from movie")
	checkErr(err)

	for rows.Next() {
		movie := movie{}
		err = rows.Scan(&movie.id, &movie.name, &movie.publishTime)
		checkErr(err)
		fmt.Println(movie.id, movie.name, movie.publishTime)
	}
	rows.Close()
}

// 删除数据
func delete() {
	stmt, err := db.Prepare("delete from movie where id=?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
