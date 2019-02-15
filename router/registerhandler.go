package router

import (
	"examps/router-fw/router/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine){
	engine.PATCH("/sendMember", handler.SendingMemberHandler)
	engine.POST("/addMember", handler.AddMemberHandler)
}
