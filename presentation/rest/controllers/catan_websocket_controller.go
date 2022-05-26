package controllers

import "github.com/kataras/iris/v12/websocket"

type CatanWebsocketController interface {
	Namespace() string
	OnNamespaceConnected(msg websocket.Message) error
	OnNamespaceDisconnected(msg websocket.Message) error
}

func NewCatanWebsocketController() CatanWebsocketController {
	return &catanWebsocketController{}
}

type catanWebsocketController struct {
	*websocket.NSConn `stateless:"true"`
}

func (c catanWebsocketController) Namespace() string {
	return "default"
}

func (c catanWebsocketController) OnNamespaceConnected(msg websocket.Message) error {
	return nil
}

func (c catanWebsocketController) OnNamespaceDisconnected(msg websocket.Message) error {
	return nil
}
