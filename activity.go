package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

var PostNumber Post

func ConnectToMySQL() {
	db, err := sqlx.Connect("mysql", "doctor:miyamoto@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connected", db)
	DBClient = db

}

func GetInt() error {

	tx := DBClient.MustBegin()
	tx.Get(&PostNumber, `SELECT id, test_numbers, created_at FROM numbers WHERE id=1`)
	tx.Commit()
	fmt.Println("ID: ", PostNumber.ID)
	fmt.Println("Number: ", PostNumber.TestNumbers)
	return nil
}

func UpdateInt() error {
	PostNumber.TestNumbers++
	td := time.Now()
	DBClient.Exec("UPDATE numbers SET test_numbers=? WHERE id=1;", PostNumber.TestNumbers)
	fmt.Println("Updated number:", PostNumber.TestNumbers, "Created: ", PostNumber.CreatedAt, "Updated: ", td)
	return nil

}

// The sunction below are for testing the  functions without a Temporal workflow
// with the main.go in the sql folder

func GetInt2(c *gin.Context) {

	td := time.Now()
	tx := DBClient.MustBegin()
	tx.Get(&PostNumber, `SELECT id, test_numbers, created_at FROM numbers WHERE id=1`)
	tx.Commit()
	fmt.Println("ID: ", PostNumber.ID)
	fmt.Println("Number: ", PostNumber.TestNumbers)

	c.JSON(http.StatusOK, gin.H{
		//"error":    false,
		"id":             PostNumber.ID,
		"Number: ":       PostNumber.TestNumbers,
		"Updated number": PostNumber.TestNumbers,
		"Created ":       PostNumber.CreatedAt,
		"Updated ":       td,
	})
	//return nil
}

func UpdateInt2(c *gin.Context) {
	PostNumber.TestNumbers++
	td := time.Now()
	DBClient.Exec("UPDATE numbers SET test_numbers=? WHERE id=1;", PostNumber.TestNumbers)
	fmt.Println("Updated number:", PostNumber.TestNumbers, "Created: ", PostNumber.CreatedAt, "Updated: ", td)

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
	})

}
