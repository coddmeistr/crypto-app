package transport

import (
	"github.com/maxim12233/crypto-app-server/account/config"
	"github.com/maxim12233/crypto-app-server/account/endpoints"
)

func NewHttpHandler(eps endpoints.AccountEndpoint) {
	config := config.GetConfig()
	r := NewHTTPServer(eps)
	r.Run(config.GetString("server.port"))
}
