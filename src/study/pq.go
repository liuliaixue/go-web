package study

import (
	"database/sql"
	"fmt"
	"log"

	// import pq
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
	}
}

// AddNote ...
func AddNote(userID string, content string) {
	connStr := "user=timqian dbname=postgres password='' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	const query = "INSERT INTO public.note(content, create_by) values($1, $2)"
	fmt.Printf(query)
	rows, err := db.Query(query, content, userID)
	if err != nil {
		fmt.Printf("error")
		log.Fatal(err)
	}
	fmt.Println(rows)

	fmt.Println("nice")
}
