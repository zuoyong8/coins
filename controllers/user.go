package controllers

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/zuoyong8/coins/models"
)


func Register(c *gin.Context){
	username := c.Param("username")
	pwd := c.Param("pwd")

	key,result := models.GetPhraseAndSecret(pwd,6)
	user := models.Users{
		Username: username,
		Pwdsalt:  key,
		Password: result,
		CreatAt: time.Now(),
	}
	err := user.Insert()
	if err != nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"msg":	   "register ok",
			"username": username,
		})
	}
}