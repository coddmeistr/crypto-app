package ws

import (
	"encoding/json"

	app "github.com/maxim12233/crypto-app-server/crypto"
)

var UnknownRequestResponse []byte
var UnknownChannelResponse []byte
var UnknownEventResponse []byte
var UnknownInternalErrorResponse []byte

func initTypes() {
	var err error
	UnknownRequestResponse, err = json.Marshal(Response{
		Event:   "error",
		Error:   app.ErrUnknownRequest.Error(),
		Message: "Unsupported request type. Need JSON",
	})
	if err != nil {
		panic("Failed initializing types JSON respresentation")
	}

	UnknownChannelResponse, err = json.Marshal(Response{
		Event:   "error",
		Error:   app.ErrUnknownChannel.Error(),
		Message: "Given channel doesn't exitst",
	})
	if err != nil {
		panic("Failed initializing types JSON respresentation")
	}

	UnknownEventResponse, err = json.Marshal(Response{
		Event:   "error",
		Error:   app.ErrUnknownEvent.Error(),
		Message: "Given event doesn't exist",
	})
	if err != nil {
		panic("Failed initializing types JSON respresentation")
	}

	UnknownInternalErrorResponse, err = json.Marshal(Response{
		Event:   "error",
		Error:   app.ErrUnknown.Error(),
		Message: "Some error occured, check message body",
	})
	if err != nil {
		panic("Failed initializing types JSON respresentation")
	}

}

type Action struct {
	Event   Event   `json:"event,omitempty"`
	Channel Channel `json:"channel,omitempty"`
	Params  string  `json:"params,omitempty"`
}

type Response struct {
	Event   string `json:"event,omitempty"`
	Channel string `json:"channel,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
