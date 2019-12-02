package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type users struct {
	id         int64
	auth_token string
	user_name  string
}

func main() {
	hoge := "data from server"
	router := gin.Default()
	// db, err := dbInit()
	// if err != nil {
	// 	fmt.Println("DB接続エラー")
	// 	return
	// }

	router.LoadHTMLGlob("view/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": hoge})
	})
	// router.GET("users", "users.html", gin, H{selectUser(db)})
	router.Run()
		fmt.Println("db acsess seccess")
}

// func dbInit() (sql.DB, error) {
// 	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=takahashi password=password dbname=postgres sslmode=disable")
// 	defer db.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 		return db, err
// 	}
// 	fmt.Println("db acsess seccess")
// 	return db, nil
// }

func selectUser(db sql.DB) users {
	var users users
	// SELECT
	// rows, err := db.Query("SELECT * FROM users;")

	// if err != nil {
	//     fmt.Println(err)
	// }
	return users
}
