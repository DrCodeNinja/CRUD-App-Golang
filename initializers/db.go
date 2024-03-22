package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func ConnectToDB() {
	LoadEnvVariables()
    // Connect to the MySQL database
    db, err := sql.Open("mysql", os.Getenv("DB_URL"))
    if err != nil {
        panic(err.Error())
    }
	DB=db
    //defer db.Close()

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Successfully connected to the database")

    // Now you can perform database operations using db
}
func CloseConnection() {
	DB.Close()
	fmt.Println("DB Connection Closed")
}