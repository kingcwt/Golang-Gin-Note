package routers

import (
	"golang-gin-note/controllers/current"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/default")
	apiRouters.GET("/user/", current.UserIndex)
	apiRouters.POST("/user/add", current.UserAdd)
	apiRouters.PUT("/user/put", current.UserPut)
	apiRouters.DELETE("/user/delete", current.UserDelete)
}
