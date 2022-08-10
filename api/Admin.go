package api

import (
	"database_project/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(context *gin.Context) {
	if auth.IsAdmin() {
		fmt.Println(context.Request.Method)
		context.Request.Method = "GET"
		fmt.Println(context.Request.Method)
		//context.HTML(http.StatusOK, "Admin.gohtml", gin.H{})
		context.HTML(http.StatusOK, "Admin.gohtml", gin.H{
			"username": auth.GetUser(),
		})
		context.Redirect(http.StatusSeeOther, "")
	} else {
		auth.Enter(context)
	}
}
