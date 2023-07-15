package routers

import (
	"golang-gin-note/controllers/api"
	"golang-gin-note/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api", middlewares.InitMiddleware)

	apiRouters.GET("/unix", api.UserGetIndex)
	// apiRouters.Use(middlewares.InitMiddleware)
	apiRouters.GET("/user", api.UserIndex)
	apiRouters.POST("/user/add", api.UserAdd)
	apiRouters.PUT("/user/put", api.UserPut)
	apiRouters.DELETE("/user/delete", api.UserDelete)
}
