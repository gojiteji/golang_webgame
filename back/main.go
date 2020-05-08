package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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


	// 接続
	conn, err := redis.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()


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
		name:=ctx.PostForm("text")
		if  name==""{
			name="hostgame"
		}
		//hostをサーバーに保存
		conn.Do("HSET", id, "members", "1")
		conn.Do("HSET", id, "host_id", userid)
		conn.Do("HSET", id, "host_name", name)

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

		id:=ctx.PostForm("id")
		name:=ctx.PostForm("text")
		if  name==""{
			name="guest"
		}
		//ルームの存在確認
		data, _ := redis.Bytes(conn.Do("HGET", id,"host_id"))
		if data != nil {
			//ルームが満室か確認
			members,_:=redis.Int(conn.Do("HGET", id, "members"))
			if(members<6){
				conn.Do("HSET", id, "members", members+1)
				conn.Do("HSET", id, "guest"+string(members)+"_id", userid)
				conn.Do("HSET", id, "guest"+string(members)+"_name", name)
			}
			ctx.Redirect(302, "/Join")
			ctx.HTML(200, "index.html", gin.H{"page":2,"message":"room id not found!"})
		}
	})

	//Room page
	router.GET("/play/:session_id", func(ctx *gin.Context) {

		//セッションidが存在するか
		//available:=true

		//クッキーからユーザを取得

		//session.Get("userid")
		//同一セッション内のユーザを全て取得
		//Users:= [...] User{host, host,host}

		//ctx.HTML(200, "waiting.html", gin.H{"host": host,"available":available,"Users":Users})
	})


	router.Run()
}