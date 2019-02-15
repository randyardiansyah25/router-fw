package handler

import (
	"examps/router-fw/helper"
	"examps/router-fw/helper/net/net"
	"examps/router-fw/model"
	"examps/router-fw/model/datastore"
	"examps/tcpclient/tcp"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
	"net/http"
	"os"
	"strconv"
)

func SendingMemberHandler(c *gin.Context){
	members, err := datastore.GetMembers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response_code" : "01",
			"response_message": "failed : "+err.Error(),
		})
		return
	}

	for no, member := range members.Data {
		_ = glg.Log("Send record : ", no+1)
		go send(member)
	}
}

func send(member model.Member){
	builder := helper.NewMessageBuilder()
	builder.Add("ID", string(member.ID))
	builder.Add("NAME", member.Name)
	builder.Add("ADDRESS", member.Address)
	builder.Add("EMAIL", member.Email)

	msg := builder.Compose()
	msg = builder.SetHeader(msg)

	sto := os.Getenv("READ_TIMEOUT")
	to, _:= strconv.ParseInt(sto, 10, 64)

	sport := os.Getenv("HOST_PORT")
	port, _:= strconv.Atoi(sport)
	client := tcp.NewTCPClient(os.Getenv("HOST_ADDR"), port, to)
	st := client.Send(msg)
	if st.Code == net.CONNOK {
		builder.Parse(st.Message)
		_ = glg.Log("RESPONSE : ",builder.Get("REPSONSE_CODE"), ">>", builder.Get("RESPONSE_MESSAGE"))
	}else if st.Code == net.ECONNTIMEOUT {
		_ = glg.Log("TIMEOUT ")
	}else {
		_ = glg.Log("Failed")
	}
}

