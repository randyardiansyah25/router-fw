package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
	"os"
)

func Start(){
	router := gin.Default()
	router.Use(gin.Recovery())
	RegisterHandler(router)

	port := os.Getenv("HTTP_LISTENER_PORT")
	_ = glg.Log("Listening at : ", port)
	err := router.Run(":"+port)
	if err != nil {
		_ = glg.Error(err.Error())
	}
}
