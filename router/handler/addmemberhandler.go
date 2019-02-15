package handler

import (
	"examps/router-fw/model"
	"examps/router-fw/model/datastore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMemberHandler(c *gin.Context){
	_ = c.Request.ParseForm()

	member := model.Member{
		Name: c.Request.FormValue("member_name"),
		Address: c.Request.FormValue("member_address"),
		Email: c.Request.FormValue("member_email"),
	}

	memberId, err := datastore.AddMember(member)
	var resp = gin.H{}
	if err != nil {
		resp["response_code"] = "01"
		resp["response_message"] = "Add member failed : "+err.Error()
	}else{
		resp["response_code"] = "00"
		resp["response_message"] = fmt.Sprint("Add member succes, your member id : ", memberId)
	}

	c.JSON(http.StatusOK, resp)
}
