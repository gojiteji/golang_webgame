package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Info struct {
	gorm.Model
	Players   int
	Page int
}

type Host struct {
	gorm.Model
	id   int
	name string
}

func Generate_title(ctx *gin.Context) {
	var info Info
	info.Players=100
	info.Page=0
	ctx.HTML(200, "index.html", gin.H{"info":info})
}

func Createroom(ctx *gin.Context){
	//redisにid登録
	var host Host
	host.id=777
	host.name="koki"
	ctx.HTML(200, "index.html", gin.H{"host":host})
}

func Joinroom(ctx *gin.Context){
	//redisにid登録
	var host Host
	host.id=777
	host.name="koki"
	ctx.HTML(200, "index.html", gin.H{"host":host})
}