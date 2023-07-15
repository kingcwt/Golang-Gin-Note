package current

import "github.com/gin-gonic/gin"

func UserIndex(r *gin.Context) {
	r.String(200, "get current-user")
}

func UserAdd(r *gin.Context) {
	r.String(200, "post current-user/add")
}

func UserPut(r *gin.Context) {
	r.String(200, "put current-user/put")
}
func UserDelete(r *gin.Context) {
	r.String(200, "delete current-user/delete")
}
