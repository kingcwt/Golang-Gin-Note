package users

import (
	"fmt"
	"golang-gin-note/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func find() []models.Users {
	userList := []models.Users{}
	// models.DB.First(&userList)
	models.DB.Where("age>20").Find(&userList)
	return userList
}

func create() string {
	user := models.Users{
		Username: "李四",
		Age:      30,
		Email:    "lisi@qq.com",
		AddTime:  time.Now().Unix(),
	}

	d := models.DB.Create(&user)
	fmt.Println(d)
	return "创建成功"
}

func update() string {

	user := models.Users{}
	models.DB.Where("id = ?", 2).Find(&user)
	user.Username = "张三在写代码"
	user.Email = "zzzz@qq.com"
	user.AddTime = time.Now().Unix()
	models.DB.Save(&user)

	return "更新成功"
}

func update2() string {

	user := models.Users{Id: 1}
	models.DB.Find(&user)
	user.Username = "king"
	user.Email = "king@qq.com"
	models.DB.Save(&user)

	return "更新成功"
}

func update3() string {
	user := models.Users{}
	models.DB.Model(&user).Where("id = ?", 3).Updates(models.Users{Username: "李四update", Email: "lisiupdate@qq.com"})
	return "更新成功1"
}

func delete() string {
	user := models.Users{}
	models.DB.Where("id = ?", 3).Delete(&user)
	return "删除成功"
}

func (u UsersController) GetList(r *gin.Context) {

	// 查询数据库
	// data := find()
	// data := create()
	data := delete()
	r.JSON(200, gin.H{
		"data": data,
	})
}
