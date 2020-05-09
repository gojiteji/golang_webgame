package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.Static("/images", "./images")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	// 接続
	router.LoadHTMLGlob("*.html")

	//title
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"page":0})


	})


	router.Run()
}