package routers

import (
	"golang-gin-note/controllers/article"

	"github.com/gin-gonic/gin"
)

func ArticleRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/article")
	apiRouters.GET("/list", article.ArticleController{}.Index)
}
