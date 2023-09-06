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

var db *sql.DB

func main() {
	var err error

	// connStr := "postgres://sriram:123456@localhost/mydb?sslmode=disable"
	connStr := "postgres://postgres:123456@localhost/mydb?sslmode=disable"

	// connStr := "user=postgres host=localhost dbname=mydb sslmode=disable"

	// for opening connection to local postgres server
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer fmt.Println("Closing the database connection")

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The Database is connected")

	// CreateTable()

	// Insert()

	// Update()

	QueryUsers()
}

func QueryUsers() {
	var (
		userid    int
		username  string
		password  string
		createdAt time.Time
		updatedAt time.Time
	)

	query := `SELECT userid, username, password, created_at , updated_at FROM users WHERE userid = $1`

	var id int
	fmt.Print("Enter the userId for printing it's userdata : ")
	fmt.Scan(&id)

	err := db.QueryRow(query, id).Scan(&userid, &username, &password, &createdAt, &updatedAt)

	CheckError(err)

	fmt.Printf("userId : %v\nuserName : %v\npassword : %v\ncreatedAt : %v\nupdated_at : %v\n\n", userid, username, password, createdAt.Format(time.RFC850), updatedAt.Format(time.RFC850))
}

func CreateTable() {

	query := `
	CREATE TABLE users (
	    userId SERIAL PRIMARY KEY,
	    username TEXT NOT NULL,
	    password TEXT NOT NULL,
	    created_at TIMESTAMP
	);`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err := db.Exec(query)

	CheckError(err)
}

func Insert() {

	// inserting users into the database now
	// sample data
	// username := "johndoe"
	// password := "secret"

	createdAt := time.Now()

	var username string
	var password string

	fmt.Print("Enter username : ")
	fmt.Scan(&username)

	fmt.Print("Enter password : ")
	fmt.Scan(&password)

	fmt.Printf("username : %v\npassword : %v\ncreated_at : %v\n", username, password, createdAt.Format(time.RFC850))

	var insertcmd string

	insertcmd = `insert into "users"("username", "password","created_at","updated_at") values( $1 , $2 , $3 , $4 )`
	// fmt.Println(insertcmd)
	_, err := db.Exec(insertcmd, username, password, createdAt, createdAt)
	// result, err := db.Exec(insertcmd, username, password, createdAt, createdAt)
	// result, err := db.Exec(`INSERT INTO users ( username , password , created_at ) VALUES ( ? , ? , ?)`, username, password, createdAt)

	CheckError(err)

	// fmt.Println("result", result)
	// userID, err := result.LastInsertId()

	// fmt.Println("userId:", userID)
}

func Update() {
	var userId int
	var username, password string
	updated_at := time.Now()

	fmt.Print("Enter the userId whose data we have to update : ")
	fmt.Scan(&userId)

	fmt.Print("Enter username : ")
	fmt.Scan(&username)

	fmt.Print("Enter password : ")
	fmt.Scan(&password)

	fmt.Printf("username : %v\npassword : %v\nupdated_at : %v\n", username, password, updated_at.Format(time.RFC850))

	updateStmt := `update "users" set "username"=$1, "password"=$2 , "updated_at"=$3 where "userid"=$4`
	_, e := db.Exec(updateStmt, username, password, updated_at, userId)
	CheckError(e)
}
