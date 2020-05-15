package main

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"math/rand"
	"strconv"
	"strings"
	"time"
)
func arrayToString(A []float64, delim string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.FormatFloat(A[i],'f',-1,64))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}
	return buffer.String()
}

func enemygenerator()  string{
	id:=float64(rand.Intn(100))
	if float64(rand.Intn(10))/10>0.5 {
		//上から出現
		x := 100+float64(rand.Intn(10))/10*380
		y := 0.0
		vx:=float64(rand.Intn(10))/10-0.5
		vy:=float64(rand.Intn(10))/10
		arr := arrayToString([]float64{id,x,y,vx,vy}," ")
		return "[host "+ arr+"]"

	}else{
		//下から出現
		x := float64(rand.Intn(10))/10*380
		y := 285.0
		vx:=float64(rand.Intn(10))/10-0.5
		vy:=-float64(rand.Intn(10))/10
		arr := arrayToString([]float64{id,x,y,vx,vy}," ")
		return "[host "+ arr+"]"
	}

}
func main() {
	rand.Seed(time.Now().UnixNano())


	router := gin.Default()
	router.Static("/images", "./images/")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	// 接続
	router.LoadHTMLGlob("*.html")

	mrouter := melody.New() //melodyのルーター

	//cookie
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//title

	router.GET("/", func(ctx *gin.Context){
				ctx.HTML(200, "index.html", gin.H{"page": 0, "error": -1})
	})

	router.POST("/", func(ctx *gin.Context){

		room:=ctx.PostForm("room")
		name:=ctx.PostForm("name")
		host:=ctx.PostForm("host")

		if(room==""||name==""){
		ctx.HTML(200, "index.html", gin.H{"page": 0, "error": -1})
		}else {
			room := ctx.PostForm("room")
			name := ctx.PostForm("name")
			if room == "" {
				ctx.HTML(200, "index.html", gin.H{"page": 0, "error": 1})
			} else if name == "" {
				ctx.HTML(200, "index.html", gin.H{"page": 0, "error": 2})
			} else {
				if host=="2"{
					session := sessions.Default(ctx)
					session.Set("username", name)
					session.Set("ishost", "1")
					session.Save()
					//url/:idに転送
					ctx.Redirect(302, "/room/"+room)
				}else if host=="1"{
					session:= sessions.Default(ctx)
					session.Set("username", name)
					session.Set("ishost", "0")
					session.Save()

					//ルーム確認
					//ctx.HTML(200, "index.html", gin.H{"page": 0, "error": -5})
					//ルーム入室
					ctx.Redirect(302, "/room/"+room)
					//broadcastでスタート
					//ctx.HTML(200, "index.html", gin.H{"page": 0,"error":5})
				}
			}
		}
	})

	router.GET("/room/:id", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		username := session.Get("username")
		ishost:=session.Get("ishost")
		id:= ctx.Param("id")
		ctx.HTML(200, "index.html", gin.H{"page": 1, "error": -1,"name":username,"ishost":ishost,"roomid":id})
	})

	router.GET("/ws/:id", func(ctx *gin.Context) {
		mrouter.HandleRequest(ctx.Writer, ctx.Request)
	})

	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
		params := strings.Split(string(msg), " ")
		if(params[0]=="newuser") {
			mrouter.Broadcast([]byte(fmt.Sprintf("%s", params)))
		}
		if(params[0]=="gethostname"||params[0]=="hostname") {
			mrouter.BroadcastOthers([]byte(fmt.Sprintf("%s", params)), s)
		}

		if(params[0]=="locationupdate"||params[0]=="getenemy") {
			mrouter.BroadcastOthers([]byte(fmt.Sprintf("%s", params)), s)
		}

		if(params[0]=="host"){
			mrouter.Broadcast([]byte(fmt.Sprintf("%s", enemygenerator())))
		}
		if(params[0]=="sc"){
				mrouter.Broadcast([]byte(fmt.Sprintf("%s", params)))
		}
		if(params[0]=="time"){
			mrouter.Broadcast([]byte(fmt.Sprintf("%s", params)))
		}

	})

	router.Run()

}