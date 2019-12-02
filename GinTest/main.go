package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	Id         int64
	AuthToken string
	UserName  string
}

func main() {
	hoge := "data from server"
	router := gin.Default()
	db, err := dbInit()
	if err != nil {
		fmt.Println("DB接続エラー")
		return
	}
	router.LoadHTMLGlob("view/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": hoge})
	})
	router.GET("users", func(ctx *gin.Context) {
		ctx.HTML(200, "users.html", gin.H{"User": selectUser(db)})
	})
	router.Run()
	defer db.Close()
	fmt.Println("db acsess seccess")
}

func dbInit() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", "host=127.0.0.1 port=5432 user=takahashi password=password dbname=postgres sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("db acsess seccess")
	return db, nil
}

func selectUser(db *sql.DB) []User {
	// SELECT
	rows, err := db.Query("SELECT * FROM users;")
	
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	var users []User
	for rows.Next() {
			var user User
			rows.Scan(&user.Id, &user.AuthToken, &user.UserName)
			fmt.Println(user)
			users = append(users, User{
						Id:     user.Id,
						AuthToken:     user.AuthToken,
						UserName:     user.UserName,
					})
	}
	fmt.Printf("%v", users)
	return users
}
