package controllers

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/zuoyong8/coins/models"
)


func Register(c *gin.Context){
	username := c.PostForm("username")
	pwd := c.PostForm("pwd")
	if len(pwd)<6 || len(pwd)>20{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": "密码小于6位或大于20位",
		})
		return
	}
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