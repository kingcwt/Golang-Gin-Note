package admin

import (
	"golang-gin-note/controllers/common"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.BaseController
}

// todo: 把方法绑定到结构体上
func (con UserController) Index(r *gin.Context) {
	// r.String(200, "get admin-user")
	con.Success(r)
}

func (con UserController) Add(r *gin.Context) {
	r.String(200, "post admin-user/add")
}

func (con UserController) Put(r *gin.Context) {
	r.String(200, "put admin-user/put")
}
func (con UserController) Delete(r *gin.Context) {
	r.String(200, "delete admin-user/delete")
}
