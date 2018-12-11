package main

import (
	"strconv"
	"./ethereum"
	"./bitcoin"
	"github.com/gin-gonic/gin"
	 "./common"
	 "./usdt"
)

func main(){
	
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/bitcoin/validateaddress/:address", func(c *gin.Context) {
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
	})
	router.GET("/bitcoin/getbalance", func(c *gin.Context) {
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
	})

	router.GET("/bitcoin/getnewaddress", func(c *gin.Context) {
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
	})
	router.GET("/bitcoin/getblockcount", func(c *gin.Context) {
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
	})
	router.GET("/bitcoin/listaccounts", func(c *gin.Context) {
		accounts,err := bitcoin.ListAccounts()
		if err!=nil{
			c.JSON(500, gin.H{
				"status":  "failure",
				"err": err,
			})
		}else{
			c.JSON(200, gin.H{
				"status":  "success",
				"blockcount": accounts,
			})
		}
	})
	router.GET("/bitcoin/listtransactions", func(c *gin.Context) {
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
	})
	router.GET("/bitcoin/listaddressgroupings", func(c *gin.Context) {
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
	})
	router.GET("/bitcoin/gettransaction/:txid", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/gethavebalancewithaddress", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/gettransactioncount/:address", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/getgasprice", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/newblockfilter", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/getfilterchanges/:filtercode", func(c *gin.Context) {
		filterCode := c.Param("filtercode")
		result,err := ethereum.GetFilterChanges(filterCode)
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
	})
	router.GET("/ethereum/getblockbyhash/:hash", func(c *gin.Context) {
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
	})
	router.GET("/ethereum/gettransactionbyhash/:hash", func(c *gin.Context) {
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
	})
	router.GET("/usdt/getwalletaddressbalances", func(c *gin.Context) {
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
	})
	router.GET("/usdt/listtransactions/:address", func(c *gin.Context) {
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
	})
	router.GET("/usdt/getinfo", func(c *gin.Context) {
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
	})
	router.GET("/usdt/gettransaction/:txid", func(c *gin.Context) {
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
	})
	router.GET("/usdt/listblocktransactions/:index", func(c *gin.Context) {
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
	})
	result,err := ethereum.GetFilterChanges("0xedb9e72fb09baa2bce8a7408b70a832b")
	if err == nil{
		sss := result[0].BlockNumber
		if len(sss)>0{

		}
	}
}