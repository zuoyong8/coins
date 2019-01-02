package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/cihub/seelog"

	"github.com/zuoyong8/coins/models"
	"github.com/zuoyong8/coins/controllers"
)


func main(){
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("log/config.xml")  
    if err != nil{  
        log.Error("parse config.xml error")  
	} 
	log.ReplaceLogger(logger)
	db, err := models.InitDB()
	if err != nil {
		log.Error(err)
		return
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/user/register",controllers.Register)

	authMiddleware,err := controllers.JwtAuth()
	if err != nil {
		log.Error(err)
		return
	}
    router.POST("/user/login", authMiddleware.LoginHandler)

	auth := router.Group("/api")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//bitcoin
		auth.GET("/bitcoin/validateaddress/:address", controllers.ValidateAddress)
	    auth.GET("/bitcoin/getbalance", controllers.GetBalance)
	    auth.GET("/bitcoin/getnewaddress", controllers.GetNewAddress)
		auth.GET("/bitcoin/getblockcount",controllers.GetBlockCount) 
		auth.GET("/bitcoin/getconnectioncount",controllers.GetConnectionCount) 
		auth.GET("/bitcoin/dumpprivkey/:address", controllers.DumpPrivkey)
	    auth.GET("/bitcoin/listaccounts", controllers.ListAccounts)
	    auth.GET("/bitcoin/listtransactions",controllers.ListTransactions) 
	    auth.GET("/bitcoin/listaddressgroupings", controllers.ListAddressGroupings)
		auth.GET("/bitcoin/gettransaction/:txid", controllers.GetTransaction)

		//ethereum
		auth.GET("/ethereum/gethavebalancewithaddress", controllers.GetHaveBalanceWithAddress)
	    auth.GET("/ethereum/gettransactioncount/:address", controllers.GetTransactionCount)
	    auth.GET("/ethereum/getgasprice", controllers.GetGasPrice)
	    auth.GET("/ethereum/newblockfilter",controllers.NewBlockFilter)
	    auth.GET("/ethereum/getfilterchanges", controllers.GetFilterChanges)
	    auth.GET("/ethereum/getblockbyhash/:hash", controllers.GetBlockByHash)
	    auth.GET("/ethereum/gettransactionbyhash/:hash", controllers.GetTransactionByHash)
		//usdt
		auth.GET("/usdt/getwalletaddressbalances", controllers.GetWalletaddressBalances)
	    auth.GET("/usdt/listtransactions/:address",controllers.UsdtListTransactions)
		auth.GET("/usdt/getinfo", controllers.Getinfo)
		auth.GET("/usdt/gettransaction/:txid", controllers.UsdtGetTransaction)
		auth.GET("/usdt/listblocktransactions/:index",controllers.ListBlockTransactions)
	}
	router.Run(":8080")
}