package main

import (
	"fmt"
	"gopkg.in/olahol/melody.v1"
	"strings"
)

func main(){
	mrouter := melody.New() //melodyのルーター

	//クライアント接続時
	mrouter.HandleConnect(func(s *melody.Session) {
	})

	// クライアント切断時
	mrouter.HandleDisconnect(func(s *melody.Session) {

	})

	// クライアントから受信時に動作

	}
}
