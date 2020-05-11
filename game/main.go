package main

import (
	"bytes"
	"math/rand"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
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
	//title
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"page":0})


	})

	router.GET("/ws", func(ctx *gin.Context) {
		mrouter.HandleRequest(ctx.Writer, ctx.Request)
	})

	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
		params := strings.Split(string(msg), " ")
		if(params[0]=="locationupdate"||params[0]=="getenemy") {
			mrouter.BroadcastOthers([]byte(fmt.Sprintf("%s", params)), s)
		}

		if(params[0]=="host"){
			mrouter.Broadcast([]byte(fmt.Sprintf("%s", enemygenerator())))
		}
	})

	router.Run()

}