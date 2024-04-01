package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Data struct {
	id   int
	name string
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	//result, err := db.Exec("insert into data values(4, 'pqrst')")
	//checkError(err)
	//rowsAffected, err := result.RowsAffected()
	//checkError(err)
	//fmt.Println("rowsAffected: ", rowsAffected)

	//lastInsertedId, err := result.LastInsertId()
	//checkError(err)
	//fmt.Println("lastInsertedId: ", lastInsertedId)

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}

}
