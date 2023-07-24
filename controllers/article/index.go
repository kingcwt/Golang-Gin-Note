package article

import (
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func findArticle(c *gin.Context) {
	// 查询文章 获取文章对应的分类
	articleList := []models.Article{}
	models.DB.Preload("ArticleCate").Find(&articleList)
	c.JSON(200, gin.H{"data": articleList})
}

func findArticleCate(c *gin.Context) {
	// 查询文章分类下包含的文章
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200, gin.H{"data": articleCateList})
}

func (con ArticleController) Index(c *gin.Context) {
	// findArticle(c)
	findArticleCate(c)
}
