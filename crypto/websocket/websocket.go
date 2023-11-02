package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewWebsocketServer(r *gin.Engine) *gin.Engine {

	initTypes()
	hub := newHub()
	go hub.run()

	r.GET("/ws", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}))

	return r
}
