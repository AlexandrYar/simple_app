package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	conString := "host=localhost port=5432 user=postgres password=10unibos dbname=users sslmode=disable"

	conn, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func FindLogin(conn *sql.DB, loginUser, passwordUser string) (string, bool) {
	rows, err := conn.Query(`select "login" from user_info where login =$1`, loginUser)
	var message string = " "
	log.Print(">" + loginUser + "<>" + passwordUser + "<")
	if loginUser == "" && passwordUser == "" {
		return message, false
	}
	if err != nil {
		message = fmt.Sprint("\nНе найден пользователь с логином: ", loginUser)
		log.Print(err)
		return message, false
	}
	var login string
	for rows.Next() {
		rows.Scan(&login)
	}

	rows, err = conn.Query(`select "password" from user_info where password =$1`, passwordUser)
	var password string
	if (err != nil) && (login == loginUser) {
		message = "Неверный пароль"
		log.Print(err)
		return message, false
	}
	for rows.Next() {
		rows.Scan(&password)
	}

	if (login == loginUser) && (password == passwordUser) {
		message = fmt.Sprint("Пользователь " + loginUser + " найден")
		log.Print("User " + loginUser + " is login now!")
		return message, true
	} else {
		message = "Неверное имя пользователя или пароль"
		return message, false
	}
}

func Register(conn *sql.DB, login, password, first_name, second_name, email, date_of_birth string) {
	sqlStatement := `insert into user_info ( id, login, password, first_name, last_name, email, date_of_birth) values ($1, $2, $3, $4, $5, $6, $7)`
	rows, err := conn.Query(`SELECT COUNT(*) FROM user_info`)
	if err != nil {
		log.Fatal(err)
	}
	var id int

	for rows.Next() {
		rows.Scan(&id)
	}
	_, e := conn.Exec(sqlStatement, id+1, login, password, first_name, second_name, email, date_of_birth)
	fmt.Println(e, "\nNew user!!")
}
