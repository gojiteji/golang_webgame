package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"

	//"./pages"
	"github.com/jinzhu/gorm"
)


type User struct{
	 gorm.Model
	 Id string
	 Name string
	 IsHost bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString1(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}





func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	//title
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"page":0})
	})

	//CreateRoom
	router.POST("/Create", func(ctx *gin.Context) {
		ctx.Redirect(302, "/")
		ctx.HTML(200, "index.html", gin.H{"page":1,"Name":"gamehost"})
	})

	//Created
	router.POST("/Created", func(ctx *gin.Context) {
		//roomidを生成
		id :=RandString1(20)
		var host User
		host.Id=ctx.Param("session_id")
		host.Name=ctx.PostForm("text")
		host.IsHost=true
		//hostをサーバーに保存
		ctx.Redirect(302, "/play/"+string(id))
	})

	//JoinGame
	router.POST("/Join", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"page":2,"Name":"guest"})
	})

	router.POST("/Joined", func(ctx *gin.Context) {
		//guestをサーバーに保存
		id:=ctx.PostForm("id")
		ctx.Redirect(302, "/play/"+id)
	})

	//Room page
	router.GET("/play/:session_id", func(ctx *gin.Context) {

		//使えるサーバーかどうか(ログイン済みか?参加ユーザか?)
		available:=true

		//guest
		var guest User
		guest.Id=ctx.Param("id")
		guest.Name=ctx.PostForm("text")
		guest.IsHost=false

		//host
		var host User
		host.Id=ctx.Param("session_id")
		host.Name="hostgame"
		host.IsHost=true

		//同一セッション内のユーザを全て取得
		Users:= [...] User{host, host,host}

		ctx.HTML(200, "waiting.html", gin.H{"host": host,"available":available,"Users":Users})
	})
	router.Run()
}