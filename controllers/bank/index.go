package bank

import (
	"golang-gin-note/controllers/common"
	"golang-gin-note/models"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	common.BaseController
}

func (bank BankController) Index(c *gin.Context) {

	// c.JSON(200, gin.H{
	// 	"code": 200,
	// 	"data": "ok",
	// })
	// 开启事务
	tx := models.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance + 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		bank.Error(c)
		return
	}

	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance - 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		bank.Error(c)
		return
	}
	tx.Commit()
	bank.Success(c)

	c.JSON(200, gin.H{
		"u1": u1,
		"u2": u2,
	})
}
