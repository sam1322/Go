package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// var err error

	// connStr := "postgres://sriram:123456@localhost/mydb?sslmode=disable"
	connStr := "postgres://postgres:123456@localhost/mydb?sslmode=disable"

	// connStr := "user=postgres host=localhost dbname=mydb sslmode=disable"

	// for opening connection to local postgres server
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The Database is connected")

	// query := `
	// CREATE TABLE users (
	//     userId SERIAL PRIMARY KEY,
	//     username TEXT NOT NULL,
	//     password TEXT NOT NULL,
	//     created_at TIMESTAMP
	// );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	// _, err = db.Exec(query)
	// if err != nil {
	// 	panic(err)
	// }

	// inserting users into the database now

	// sample data

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	fmt.Println(username, password, createdAt)

	var insertcmd string

	insertcmd = `insert into "users"("username", "password","created_at") values($1, $2,$3)`
	fmt.Println(insertcmd)
	result, err := db.Exec(insertcmd, username, password, createdAt)
	// result, err := db.Exec(`INSERT INTO users ( username , password , created_at ) VALUES ( ? , ? , ?)`, username, password, createdAt)

	CheckError(err)

	fmt.Println("result", result)
	// userID, err := result.LastInsertId()

	// fmt.Println("userId:", userID)

}
