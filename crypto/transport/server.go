package transport

import (
	"github.com/maxim12233/crypto-app-server/crypto/config"
	"github.com/maxim12233/crypto-app-server/crypto/endpoints"
	ws "github.com/maxim12233/crypto-app-server/crypto/websocket"
)

func NewHttpHandler(eps endpoints.CryptoEndpoint) {
	config := config.GetConfig()
	r := NewHTTPServer(eps)
	r = ws.NewWebsocketServer(r)
	r.Run(config.GetString("server.port"))
}
