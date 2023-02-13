package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connect to database
	dbConnect, err := sql.Open("pgx", "host=localhost port=5432 dbname=database_connect user=postgres password=")

	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	defer dbConnect.Close()

	log.Println("Connected to database")

	// test databse connection
	err = dbConnect.Ping()
	if err != nil {
		log.Fatal("Cannot ping database")
	}
	log.Println("Pinged database")

	// get rows from database
	err = getAllRows(dbConnect)
	if err != nil {
		log.Fatal(err)
	}

	// Insert into table
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = dbConnect.Exec(query, "Ewomazino", "Atu")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row")

	// Get updated rows from table
	err = getAllRows(dbConnect)
	if err != nil {
		log.Fatal(err)
	}

	// Update a row
	updateUser := `update users set first_name = $1 where id = $2`
	_, err = dbConnect.Exec(updateUser, "Ewoma", 4)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated a user")

	// Get updated rows from table
	err = getAllRows(dbConnect)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(dbConnect *sql.DB) error {
	rows, err := dbConnect.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var first_name, last_name string
	var id int

	for rows.Next() {
		rowsError := rows.Scan(&id, &first_name, &last_name)
		if rowsError != nil {
			log.Println(rowsError)
			return rowsError
		}
		fmt.Println("Recored is ", id, first_name, last_name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("---------------")

	return nil
}
