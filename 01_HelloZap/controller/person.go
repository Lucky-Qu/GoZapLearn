package controller

import (
	"GoZapLearn/01_HelloZap/db/model"
	"GoZapLearn/01_HelloZap/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPerson(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if person, err := service.GetPerson(id); err != nil {
		panic(err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "查询成功",
			"data": person,
		})
	}
}
func CreatePerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBind(&person); err != nil {
		panic(err)
		return
	}
	if err := service.CreatePerson(&person); err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}
func UpdatePerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBind(&person); err != nil {
		panic(err)
		return
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err = service.UpdatePerson(id, &person); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
func DeletePerson(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err = service.DeletePerson(id); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
