package navs

import (
	"golang-gin-note/models"
	"time"

	"github.com/gin-gonic/gin"
)

type NavsController struct{}

// 临时测试
type NavJson struct {
	Id    int32  `gorm:"primary_key"`
	Title string `json:"title"`
}

func (NavJson) TableName() string {
	return "navs"
}

func find(c *gin.Context) {
	// navResult := []models.Navs{}
	// AND
	// models.DB.Where("id > ? AND id < ?", 1, 9).Find(&navResult)

	// IN
	// models.DB.Where("id in (?) ", []int{3, 5}).Find(&navResult)

	// like 查询标题里面 包含代码的内容
	// models.DB.Where("title like ? ", "%代码%").Find(&navResult)

	// between 查询id在 5-7之间
	// models.DB.Where("id between ? and ?", 5, 7).Find(&navResult)

	// or 查询id = 3 or id = 6
	// models.DB.Where("id = ? or id = ?", 3, 6).Find(&navResult)
	// models.DB.Where("id = ?", 2).Or("id = ?", 6).Or("id = ?", 3).Find(&navResult)

	// select 指定返回的字段
	// navJsonResult := []NavJson{}
	// models.DB.Select("id,title").Find(&navJsonResult)

	// order 排序 limit offset   // desc 降序  // limit 2 展示降序前两条  // offset 2 跳过前两条数据
	// models.DB.Order("id desc").Offset(2).Limit(2).Find(&navResult)

	// count 统计数量
	// var num int64
	// models.DB.Find(&navResult).Count(&num)
	// c.JSON(200, gin.H{
	// 	"data":  navResult,
	// 	"count": num,
	// })

}

func nativeSqlOperat(c *gin.Context) {
	navResult := []models.Navs{}

	// --- 使用原生sql ---
	// todo Exec 执行sql语句 Raw 查询sql语句

	// 删除一条数据
	// models.DB.Exec("delete from navs where id = ?", 4)

	// 修改数据
	// models.DB.Exec("update navs set title = ? where id = ?", "id1的在大半夜学习golang", 1)

	// 查询数据
	models.DB.Raw("select * from navs where id > 5").Scan(&navResult)

	// 统计数据
	var num int64
	models.DB.Raw("select count(*) from navs").Scan(&num)
	models.DB.Find(&navResult)
	c.JSON(200, gin.H{
		"data":  navResult,
		"count": num,
	})
}

func create(c *gin.Context) {
	nav := models.Navs{
		Title:   "小米在写代码",
		Url:     "http://www.baidu.com",
		Status:  7,
		AddTime: time.Now().Unix(),
	}

	models.DB.Create(&nav)
	c.JSON(200, gin.H{
		"data": nav,
		"msg":  "创建成功",
	})
}

func (n NavsController) GetList(c *gin.Context) {
	nativeSqlOperat(c)
	// find(c)
	// create(c)

}
