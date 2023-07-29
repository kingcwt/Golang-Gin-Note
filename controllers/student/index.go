package student

import (
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct{}

func (stu StudentController) Index(c *gin.Context) {
	// 获取学生信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 获取所有课 程信息
	// lessonList := []models.Lesson{}
	// models.DB.Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	// 查询学生信息的时候 展示学生选修的课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})
	// 获取学生选修的课程信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 查询张三选修了哪些课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Where("name = ?", "张三").Find(&studentList)
	// c.JSON(200, gin.H{"data": studentList})

	// 课程被哪些学生选修了
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student").Limit(2).Order("id desc").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	//查询课程被哪些学生选修的同时去掉张三
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "name != ?", "张三").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	//查询课程被哪些学生选修的同时去掉张三和李四
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "id not in (1,2)").Find(&lessonList)
	// c.JSON(200, gin.H{"data": lessonList})

	// 预加载SQL 在查询课程被哪些学生选修的同时 学生数据倒叙排列
	lessonList := []models.Lesson{}
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return models.DB.Where("id > 1").Order("id desc")
	}).Find(&lessonList)
	c.JSON(200, gin.H{"data": lessonList})
}
