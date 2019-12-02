package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("tmp/*.html")
	router.GET("/",func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.Run()

	fmt.Println("success")
}