package main

import (
	"first/initializers"
	"fmt"
	"log"
)

//The init function is automatically executed before the main function when the program starts.
func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}


func main() {
	createTable()
}
func createTable() {
	// Define the SQL statement to create the table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(255) NOT NULL,
		Body VARCHAR(500) NOT NULL
	);
	`
	
	// Execute the SQL statement to create the table
	_, err := initializers.DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	initializers.CloseConnection()
	fmt.Println("Table created successfully")
}
