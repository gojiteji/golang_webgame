package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"./pages"
)

func acces(session_id string,ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"session_id": session_id})
}


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	//title
	router.GET("/", func(ctx *gin.Context){
		pages.Generate_title(ctx)
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		ctx.Redirect(302, "/")
		pages.Createroom(ctx)
	})



	router.GET("/play/:session_id", func(ctx *gin.Context){
		session_id:= ctx.Param("session_id")
		fmt.Print("session_id is")
		fmt.Print(session_id)
		acces(session_id,ctx)
	})


	router.Run()
}