package routers

import (
	"golang-gin-note/controllers/bank"

	"github.com/gin-gonic/gin"
)

func BankRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/bank")
	apiRouters.GET("/list", bank.BankController{}.Index)
}
