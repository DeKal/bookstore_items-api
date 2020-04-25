package controllers

import "net/http"

// PingControllerInterface is an interface for itemService
type PingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

// NewPingController return new itemService
func NewPingController() PingControllerInterface {
	return &pingController{}
}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
