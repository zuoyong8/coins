package controllers

import(
	"strconv"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	// "math/big"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/zuoyong8/coins/ethereum"
	"github.com/zuoyong8/coins/common"
)

func GetCoinbaseAddress(c *gin.Context){
	address,err := ethereum.Coinbase()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"balances": address,
		})
	}
}

func GetAccouts(c *gin.Context){
	accouts,err := ethereum.GetAccouts()
	if err != nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"accouts": accouts,
		})
	}
}

//func 4 bytes + params 32byte3 不足前面补0
func EthCall(c *gin.Context){
	address := "0x230eaaf5812f6833990bc0f39085527946a043fe"
	f := solsha3.SoliditySHA3(
		solsha3.String("double(int256)"))
	// fc := "0x"+(hex.EncodeToString(f)[:8])
	fc := "0x70a08231"
	a := "000000000000000000000000230eaaf5812f6833990bc0f39085527946a043fe"
	result,err := ethereum.EthCall(address,fc+a)
	if err != nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": result,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"ethcall": hex.EncodeToString(f)[:8],
		})
	}
}

func GetBlockNumber(c *gin.Context){
	blocknumber,err := ethereum.GetBlockNumber()
	if err != nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"blocknumber": common.HexToDecimal(blocknumber),
		})
	}
}

func NetVersion(c *gin.Context){
	vesion,err := ethereum.NetVersion()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"netversion": vesion,
		})
	}
}

func Web3ClientVersion(c *gin.Context){
	vesion,err := ethereum.Web3ClientVersion()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"web3_clientversion": vesion,
		})
	}
}

func GetSyning(c *gin.Context){
	syning,err := ethereum.GetSyning()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"syning": syning,
		})
	}
}

func Web3Sha3(c *gin.Context){
	// data := c.Query("data")
	// sha3,err := ethereum.Web3Sha3(data)
	// if err!=nil{
	// 	c.JSON(500, gin.H{
	// 		"status": "failure",
	// 		"err": err,
	// 	})
	// }else{
		hash := solsha3.SoliditySHA3(
			//solsha3.Address("0x12459c951127e0c374ff9105dda097662a027093"),
			//solsha3.Uint256(big.NewInt(100)),
			solsha3.String("transfer"))
		c.JSON(200, gin.H{
			"status": "success",
			"web3sha3": hex.EncodeToString(hash),
		})
	// }
}

func GetHaveBalanceWithAddress(c *gin.Context){
	balances,err := ethereum.GetHaveBalanceWithAddress()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"balances": balances,
		})
	}
}

func EthGetBalance(c *gin.Context){
	address := c.Param("address")
	balance,err := ethereum.GetBalance(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"balances": balance,
		})
	}
}

func GetTransactionCount(c *gin.Context){
	address := c.Param("address")
	result,err := ethereum.GetTransactionCount(address)
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"transactionCount": result,
		})
	}
}

func GetGasPrice(c *gin.Context){
	gasPrice,err := ethereum.GetGasPrice()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"gasprice": common.HexToDecimal(gasPrice),
		})
	}
}

func NewBlockFilter(c *gin.Context){
	filterCode,err := ethereum.NewBlockFilter()
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
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
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"result": result,
		})
	}
}

func GetBlockByHash(c *gin.Context){
	hash := c.Param("hash")
	result,err := ethereum.GetBlockByHash(hash)
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
	   })
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"result": result,
		})
	}
}

func GetBlockByNumber(c *gin.Context){
	hash := c.Param("hash")
	result,err := ethereum.GetBlockByNumber(hash)
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
	   })
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"result": result,
		})
	}
}

func GetTransactionByHash(c *gin.Context){
	hash := c.Param("hash")
	result,err := ethereum.GetTransactionByHash(hash)
	if err!=nil{
		c.JSON(500, gin.H{
			"status": "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "success",
			"result": result,
		})
	}
}

func PersonalNewAccount(c *gin.Context){
	newAccount,err := ethereum.PersonalNewAccount()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"newaddress": newAccount,
		})
	}
}

func PersonalListWallets(c *gin.Context){
	walletsInfo,err := ethereum.PersonalListWallets()
	if err!=nil{
		c.JSON(500, gin.H{
			"status":  "failure",
			"err": err,
		})
	}else{
		c.JSON(200, gin.H{
			"status":  "success",
			"WalletsInfo": walletsInfo,
		})
	}
}

func PersonalUnlockAccount(c *gin.Context){
	address := c.Param("address")
	result,err := ethereum.PersonalUnlockAccount(address)
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