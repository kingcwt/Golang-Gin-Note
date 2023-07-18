package routers

import (
	"golang-gin-note/controllers/users"

	"github.com/gin-gonic/gin"
)

func UsersRouters(r *gin.Engine) {
	routes := r.Group("/users")
	routes.GET("/list", users.UsersController{}.GetList)
}
