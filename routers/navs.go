package routers

import (
	"golang-gin-note/controllers/navs"

	"github.com/gin-gonic/gin"
)

func NavsRouters(r *gin.Engine) {
	routes := r.Group("/navs")
	routes.GET("/list", navs.NavsController{}.GetList)
}
