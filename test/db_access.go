package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

func DBAccess() {
	// psql -d postgres -U takahashi
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=takahashi password=password dbname=postgres sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println("DB接続エラー")
		fmt.Println(err)
	}

	// INSERT
	// var empID string
	id := 3
	// number := 4444
	// err = db.QueryRow("INSERT INTO users( id, auth_token, user_name) VALUES($1,$2,$3)", id, "hoge", number).Scan(&empID)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(empID)

	fmt.Println(id)
	// SELECT rows
	rows, err := db.Query("SELECT * FROM users WHERE id=?", id)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	// SELECT row
	userId := 1
	row := db.QueryRow("SELECT id, email, user_name, password_digest FROM users WHERE id = $1", userId)
	fmt.Println(row)
	user := User{}
	err = row.Scan(&user.ID, &user.Email, &user.UserName, &user.PasswordDigest)
	if err != nil {
		fmt.Println("取得失敗")
		fmt.Println(err)
		return 
	}
	fmt.Println(user)
	return 

}
