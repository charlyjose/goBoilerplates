package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := ConfigureDB("mysql", "root", "12345", "localhost", "3306", "test", "parseTime=true")

	query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
			);`

	CreateTable(db, query)
	defer DropTable(db, "users") // optional
	InsertRow(db)
	QuerySingleRow(db)
	QueryMultipleRows(db)
	DeleteRow(db)
}

// ConfigureDB function to drop the user table
func ConfigureDB(driver, username, password, hostname, port, database, params string) (db *sql.DB) {
	// "root:12345@tcp(127.0.0.1:3306)/test?parseTime=true"
	db, err := sql.Open(driver, username+":"+password+"@tcp("+hostname+":"+port+")/"+database+"?"+params)
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return
}

// CreateTable function to create a new table
func CreateTable(db *sql.DB, query string) {
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// InsertRow function to insert a row
func InsertRow(db *sql.DB) {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	_, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	//id, err := result.LastInsertId()
	//fmt.Println(id)
}

// QuerySingleRow function to query single row
func QuerySingleRow(db *sql.DB) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"

	if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}

// QueryMultipleRows function to query multiple rows
func QueryMultipleRows(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query("SELECT id, username, password, created_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		users []user
		u     user
	)
	for rows.Next() {

		if err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", users)
}

// DeleteRow function to delete a row
func DeleteRow(db *sql.DB) {
	if _, err := db.Exec("DELETE FROM users WHERE id = ?", 1); err != nil {
		log.Fatal(err)
	}
}

// DropTable function to drop the user table
func DropTable(db *sql.DB, tablename string) {
	defer func() {
		if _, err := db.Exec("DROP TABLE users"); err != nil {
			log.Fatal(err)
		}
	}()
	return
}
