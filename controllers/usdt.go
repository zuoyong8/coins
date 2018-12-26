package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"

	"github.com/zuoyong8/coins/usdt"
)

func GetWalletaddressBalances(c *gin.Context){
	balances,err := usdt.GetWalletaddressBalances()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"balances": balances,
		})
	}
}

func UsdtListTransactions(c *gin.Context){
	address := c.Param("address")
	transactions,err := usdt.ListTransactions(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"transactions": transactions,
		})
	}
}

func Getinfo(c *gin.Context){
	nodeInfo,err := usdt.Getinfo()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"nodeinfo": nodeInfo,
		})	
	}
}

func UsdtGetTransaction(c *gin.Context){
	txid := c.Param("txid")
	transaction,err := usdt.GetTransaction(txid)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"nodeinfo": transaction,
		})
	}
}

func ListBlockTransactions(c *gin.Context){
	index,err := strconv.Atoi(c.Param("index"))
	blockTransaction,err := usdt.ListBlockTransactions(index)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"blocktransaction": blockTransaction,
		})
	}
}