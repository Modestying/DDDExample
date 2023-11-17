package server

import (
	"ddd_demo/domain/iface"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	en *gin.Engine
}

func NewGinServer() iface.IServer {
	return &GinServer{
		en: gin.Default(),
	}
}
func (g *GinServer) Start() {

}

func (g *GinServer) Exit() {

}
