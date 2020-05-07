package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Info struct {
	gorm.Model
	Players   int
	ServerStatus bool
}

func Generate_title(ctx *gin.Context) {
	var info Info
	info.Players=100
	info.ServerStatus=true
	ctx.HTML(200, "index.html", gin.H{"info":info})
}