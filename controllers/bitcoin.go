package controllers

import(
	"github.com/gin-gonic/gin"

	"github.com/zuoyong8/coins/bitcoin"
)

//
func ValidateAddress(c *gin.Context){
	address := c.Param("address")
	validInfo,err := bitcoin.ValidateAddress(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"address": validInfo.Address,
			"isvalid": validInfo.IsValid,
			"ismine":  validInfo.IsMine,
			"timestamp": validInfo.TimeStamp,
		})
	}
}

//
func GetBalance(c *gin.Context){
	balance,err := bitcoin.GetBalance()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"balance": balance,
		})
	}
}

//
func GetNewAddress(c *gin.Context){
	newAddress,err := bitcoin.GetNewaAddress()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"newaddress": newAddress,
		})
	}
}

func GetBlockCount(c *gin.Context){
	blockCount,err := bitcoin.GetBlockCount()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"blockcount": blockCount,
		})
	}
}

func GetConnectionCount(c *gin.Context){
	connectionCount,err := bitcoin.GetConnectionCount()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"blockcount": connectionCount,
		})
	}
}

func ListAccounts(c *gin.Context){
	accounts,err := bitcoin.ListAccounts()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"accounts": accounts,
		})
	}
}

func ListTransactions(c *gin.Context){
	transactions,err := bitcoin.ListTransactions()
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

func ListAddressGroupings(c *gin.Context){
	addressGroupings,err := bitcoin.ListAddressGroupings()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"addressGroupings": addressGroupings,
		})
	}
}

func GetTransaction(c *gin.Context){
	txid := c.Param("txid")
	transaction,err := bitcoin.GetTransaction(txid)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"transaction": transaction,
		})
	}
}

func DumpPrivkey(c *gin.Context){
	address := c.Param("address")
	privekey,err := bitcoin.DumpPrivkey(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"transaction": privekey,
		})
	}
}