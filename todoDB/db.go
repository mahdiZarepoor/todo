package todoDB

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// model
type Todo struct {
	Username string 
	Title string 
	Description string 
	Done bool 
	Deadline string 
}

type DSN struct {
	host string 
	port int
	user string 
	password string 
	dbname string 
}

var db *sql.DB

func init() {
	// loadin .env file in our code
	if err := godotenv.Load() ; err != nil {
		log.Println("hererererere")
		log.Fatal(err)
	}

	// getting some evironment variables we stored in .env file (postgres config)
	var dsn DSN
	dsn.host= os.Getenv("HOST")
	port := os.Getenv("PORT")
	dsn.user= os.Getenv("U") 
	dsn.password= os.Getenv("PASSWORD")
	dsn.dbname= os.Getenv("DBNAME")

	var err error
	dsn.port, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	Dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbname)
	db, err = sql.Open("postgres", Dsn)
	if err != nil {
		log.Fatal("error while opening connection to db : ", err.Error())
	}
	if err :=db.Ping(); err != nil {
		log.Fatal("don't have ping to db : ", err.Error())
	}
}

func GiveMeDb() *sql.DB{
	return db
}

func InsertTodo(user, title, description, deadline string) error {
	// the table we use for this code is named todos1 
	insertSQL := `INSERT INTO todos1 VALUES ($1,$2, $3, false,$4)`
	_, err :=db.Exec(insertSQL, user, title, description, deadline)
	if err != nil {
		return err
	}
	return nil
}

func ListAllTodos (username string) ([]Todo,error) {

	query := `SELECT * FROM todos1 WHERE username = $1`
	rows , err := db.Query(query, username) 
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var t Todo 
		if err :=rows.Scan(&t.Username, &t.Title, &t.Description, &t.Done, &t.Deadline); err != nil {
			return nil, err 
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func MakeItDone(username , title string) (int, error) {

	query := `UPDATE todos1 
				SET done = true
				WHERE username=$1 AND title=$2`
	r , err := db.Exec(query, username, title)
	if err != nil {
		return 0, err 
	}

	rowsAffected , err := r.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return int(rowsAffected), nil
}

func UndoTodo (username, title string) (int, error) {
	query := `UPDATE todos1 
				SET done = false
				WHERE username=$1 AND title=$2`
	r , err := db.Exec(query, username, title)
	if err != nil {
		return 0, err 
	}

	rowsAffected , err := r.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return int(rowsAffected), nil
}

func DeleteTodo(username, title string) (int, error){

	query := `DELETE FROM todos1 WHERE username=$1 AND title=$2`
	r, err := db.Exec(query, username, title)
	if err != nil {
		return 0, err 
	}
	rowsAff , err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAff == 0 {
		return 0 , nil
	}

	return int(rowsAff), nil
}

func DeleteAll(username string) (int, error) {
	query := `DELETE FROM todos1 WHERE username = $1`
	r, err := db.Exec(query, username)
	if err != nil {
		return 0, err
	}
	rowsAff, err :=r.RowsAffected()
	if err != nil {
		return 0 , err 
	}
	return int(rowsAff), nil
}

func TodoByTitle (username , title string) (Todo, error){

	query := `SELECT * FROM todos1 WHERE username=$1 AND title = $2`
	rows := db.QueryRow(query, username, title)

	var t Todo
	err := rows.Scan(&t.Username, &t.Title, &t.Description, &t.Done, &t.Deadline)
	if err != nil {
		// in case there is no such a todo for the user
		return Todo{} , err 
	}
	return t , nil 
}
