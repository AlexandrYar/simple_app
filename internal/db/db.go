package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ConnDb struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

type User struct {
	Login         string
	Password      string
	First_name    string
	Last_name     string
	Email         string
	Date_of_birth string
}

var NewDb = ConnDb{
	host:     "localhost",
	port:     "5432",
	user:     "postgres",
	password: "10unibos",
	dbname:   "users",
	sslmode:  "disable",
}

func (db ConnDb) Connection() *sql.DB {
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db.host, db.port, db.user, db.password, db.dbname, db.sslmode)

	conn, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func (user User) LoginUser(conn *sql.DB, loginUser, passwordUser string) (string, bool) {
	rows, err := conn.Query(`select "login" from user_info where login =$1`, loginUser)
	var message string = " "
	if loginUser == "" && passwordUser == "" {
		return message, false
	}
	if err != nil {
		message = fmt.Sprint("\nНе найден пользователь с логином: ", loginUser)
		log.Print(err)
		return message, false
	}
	for rows.Next() {
		rows.Scan(&user.Login)
	}

	rows, err = conn.Query(`select "password" from user_info where password =$1`, passwordUser)
	if (err != nil) && (user.Login == loginUser) {
		message = "Неверный пароль"
		log.Print(err)
		return message, false
	}
	for rows.Next() {
		rows.Scan(&user.Password)
	}

	if (user.Login == loginUser) && (user.Password == passwordUser) {
		message = fmt.Sprint("Пользователь " + loginUser + " найден")
		log.Print("User " + loginUser + " is login now!")
		return message, true
	} else {
		message = "Неверное имя пользователя или пароль"
		return message, false
	}
}

func (user User) Register(conn *sql.DB, login, password, first_name, second_name, email, date_of_birth string) {
	sqlStatement := `insert into user_info ( id, login, password, first_name, last_name, email, date_of_birth) values ($1, $2, $3, $4, $5, $6, $7)`
	rows, err := conn.Query(`SELECT COUNT(*) FROM user_info`)
	if err != nil {
		log.Fatal(err)
	}
	var id int

	for rows.Next() {
		rows.Scan(&id)
	}
	_, e := conn.Exec(sqlStatement, id+1, user.Login, user.Password, user.First_name, user.Last_name, user.Email, user.Date_of_birth)
	fmt.Println(e, "\nNew user!!")
}

func (user *User) Find_info(conn *sql.DB, fing_login string) {
	rows, err := conn.Query(`select "login", "first_name", "last_name", "email", "date_of_birth" from user_info where login = $1`, fing_login)
	if err != nil {
		log.Println(err, "err 1")
	}
	log.Println(rows)
	for rows.Next() {
		err = rows.Scan(&user.Login, &user.First_name, &user.Last_name, &user.Email, &user.Date_of_birth)
		if err != nil {
			log.Println(err, "err 2")
		}

		log.Println(user)

	}
}
