package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "demo"
)

var db *sql.DB

type Users struct {
	Id int
	username string
	fullname string
}

func main(){
	var err error
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connstr)
	if err != nil {
	  panic(err)
	}
	fmt.Println("Database Connected")
	defer db.Close()
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
	// fmt.Println("Insert new user")
	// err = CreateUser(Users{Id: 2000, username: "tony", fullname: "Tony Kruz"})
	// if err != nil {
	// 	panic(err)
	// }
	users, err := GetUser()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func CreateUser(user Users) error {
	query := `
		INSERT INTO users (
			id,
			username,
			fullname 
		) VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, user.Id, user.username, user.fullname)
	if err != nil {
		return err
	}
	return nil
}

func GetUser() ([]Users, error) {
	query := `
		SELECT
			id, 
			username,
			fullname
		FROM users
		ORDER BY ID OFFSET 100 LIMIT 100
	`
	rows, err := db.Query(query)
	if err != nil {
		return []Users{}, err
	}
	defer rows.Close()
	users := []Users{}
	for rows.Next() {
		var user Users
		err := rows.Scan(
			&user.Id,
			&user.username,
			&user.fullname,
		)
		if err != nil {
			return []Users{}, err
		}
		users = append(users, user)
	}
	return users, nil

}