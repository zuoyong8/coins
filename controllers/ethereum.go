package controllers

import(
	"strconv"
	"github.com/gin-gonic/gin"

	"github.com/zuoyong8/coins/ethereum"
)


func GetHaveBalanceWithAddress(c *gin.Context){
	balances,err := ethereum.GetHaveBalanceWithAddress()
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

func GetTransactionCount(c *gin.Context){
	address := c.Param("address")
	result,err := ethereum.GetTransactionCount(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"result": common.HexDec(result),
		})
	}
}

func GetGasPrice(c *gin.Context){
	gasPrice,err := ethereum.GetGasPrice()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"gasprice": common.HexDec(gasPrice),
		})
	}
}

func NewBlockFilter(c *gin.Context){
	filterCode,err := ethereum.NewBlockFilter()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"filtercode": filterCode,
		})
	}
}

func GetFilterChanges(c *gin.Context){
	filterCode := c.Query("filtercode") 
	ft,err := strconv.Atoi(c.Query("filtertype"))
	var filterType ethereum.FilterType
	switch ft{
		case 0:
			filterType = ethereum.BlockFilter
		case 1:
			filterType = ethereum.PendingTransactionFilter
		case 2:
			filterType = ethereum.BlockFilter
	}
	result,err := ethereum.GetFilterChanges(filterType,filterCode)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"result": result,
		})
	}
}

func GetBlockByHash(c *gin.Context){
	hash := c.Param("hash")
	result,err := ethereum.GetBlockByHash(hash)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
	   })
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"result":result,
		})
	}
}

func GetTransactionByHash(c *gin.Context){
	hash := c.Param("hash")
	result,err := ethereum.GetTransactionByHash(hash)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"result":result,
		})
	}
}

