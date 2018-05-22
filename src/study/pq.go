package study

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func addUser() {
	connStr := "user=timqian dbname=postgres password='' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// rows, err := db.Query("SELECT * FROM public.\"user\"")
	rows, err := db.Query("SELECT * FROM public.user")
	if err != nil {

		fmt.Println("query error")
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println(id, name)
		fmt.Println("lalala")
	}
}
