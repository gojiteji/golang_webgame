package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

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


	//JoinGame
	router.POST("/Join", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"page":2,"Name":"guest"})
	})



	//Created
	router.POST("/Created", func(ctx *gin.Context) {
		//roomidを生成
		id :=RandString1(20)
		//クッキーを作成
		userid:=RandString1(32)
		session := sessions.Default(ctx)

		if session.Get("userid") != userid {
			session.Set("userid", userid)
			session.Save()
		}
		var host User
		host.Id=ctx.Param("session_id")
		host.Name=ctx.PostForm("text")
		host.IsHost=true
		//hostをサーバーに保存
		
		ctx.Redirect(302, "/play/"+string(id))
	})

	//Joined
	router.POST("/Joined", func(ctx *gin.Context) {
		//クッキーを作成
		userid:=RandString1(32)
		session := sessions.Default(ctx)
		if session.Get("userid") != userid {
			session.Set("userid", userid)
			session.Save()
		}

		//guestをサーバーに保存
		id:=ctx.PostForm("id")
		ctx.Redirect(302, "/play/"+id)
	})

	//Room page
	router.GET("/play/:session_id", func(ctx *gin.Context) {

		//セッションidが存在するか
		//available:=true

		//クッキーからユーザを取得


		//同一セッション内のユーザを全て取得
		//Users:= [...] User{host, host,host}

		//ctx.HTML(200, "waiting.html", gin.H{"host": host,"available":available,"Users":Users})
	})


	router.Run()
}