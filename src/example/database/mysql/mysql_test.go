package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

var MysqlDB *sql.DB

func init() {
	var mysqlErr error
	var user = "root"
	var password = "123456"
	var dbName = "test"
	ip := "localhost"
	port := "3306"
	url := "@tcp(" + ip + ":" + port + ")/"
	MysqlDB, mysqlErr = sql.Open("mysql", user+":"+password+url+dbName)
	if nil != mysqlErr {
		panic(mysqlErr)
	}
	MysqlDB.SetMaxOpenConns(128)
	MysqlDB.SetMaxIdleConns(16)
	MysqlDB.SetConnMaxLifetime(2 * time.Minute)
	if mysqlErr = MysqlDB.Ping(); nil != mysqlErr {
		panic(mysqlErr)
	}
}

/***
 * 创建表
 *
 */
func TestCreateTable(t *testing.T) {
	_, err := MysqlDB.Exec(`
    CREATE TABLE IF NOT EXISTS students (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        age INT,
        gender VARCHAR(10),
        grade VARCHAR(255)
    )
`)
	if err != nil {
		fmt.Println(err)
	}
}

func TestInsertSingle(t *testing.T) {
	_, err := MysqlDB.Exec("INSERT INTO students(name, age, gender, grade) VALUES(?, ?, ?, ?)", "张三", 18, "男", "高一")
	if err != nil {
		fmt.Println(err)
	}
	defer MysqlDB.Close()
}

func TestMultiInsert(t *testing.T) {
	//同时插入多条数据
	students := [][]interface{}{
		{"李四", 19, "男", "高二"},
		{"王五", 20, "女", "高三"},
		{"赵六", 17, "女", "高一"},
	}
	stmt, err := MysqlDB.Prepare("INSERT INTO students(name, age, gender, grade) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for _, student := range students {
		_, err := stmt.Exec(student[0], student[1], student[2], student[3])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestQueryTable(t *testing.T) {
	rows, err := MysqlDB.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, age int
		var name, gender, grade string
		err := rows.Scan(&id, &name, &age, &gender, &grade)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name, age, gender, grade)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}

func TestQueryWithCondition(t *testing.T) {
	rows, err := MysqlDB.Query("SELECT * FROM students WHERE age > ? AND grade = ?", 18, "高二")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, age int
		var name, gender, grade string
		err := rows.Scan(&id, &name, &age, &gender, &grade)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name, age, gender, grade)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}

func TestMain(m *testing.M) {

	m.Run()
	defer MysqlDB.Close()
}

func TestDeleteAll(t *testing.T) {
	//删除全部数据
	_, err := MysqlDB.Exec("DELETE FROM students")
	if err != nil {
		fmt.Println(err)
	}
}

func TestDeleteTable(t *testing.T) {

	//按条件删除数据
	_, err := MysqlDB.Exec("DELETE FROM students WHERE age < ?", 18)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUpdateWithCondition(t *testing.T) {
	//按条件更新数据
	_, err := MysqlDB.Exec("UPDATE students SET age = age + 1 WHERE gender = ?", "女")
	if err != nil {
		fmt.Println(err)
	}
}

func TestUpdateTable(t *testing.T) {
	_, err := MysqlDB.Exec("UPDATE students SET age = age + 1")
	if err != nil {
		fmt.Println(err)
	}
}

func TestShowTables(t *testing.T) {
	rows, err := MysqlDB.Query("show tables")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var tableName string
	for rows.Next() {
		// 获取每一行的数据
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tableName)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}
