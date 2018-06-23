package study

import (
	sql "database/sql"
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
	db.Close()
	fmt.Println(rows)
}

// Note ...
type Note struct {
	Id       int
	Content  string
	CreateAt string
	CreateBy string
}

// GetNote nice
func GetNote() []Note {
	connStr := "user=timqian dbname=postgres password='' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	const query = "SELECT * from public.note;"
	fmt.Printf(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("error")
		log.Fatal(err)
	}
	db.Close()
	fmt.Println(rows)
	var notes []Note
	for rows.Next() {
		var id int
		var content string
		var createAt string
		var createBy string
		err := rows.Scan(&id, &content, &createAt, &createBy)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(id, content, createAt, createBy)
		note := Note{id, content, createAt, createBy}
		notes = append(notes, note)

	}
	return notes
}
