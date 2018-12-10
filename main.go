
package main


import (
	"fmt"
	"./ethereum"
	"./bitcoin"
	"github.com/gin-gonic/gin"
	 "./common"
	//  "./usdt"
)

// type MyTest struct{
// 	Out []interface{}
// 	A 	string
// }

// func COut(myout []interface{},mya string)*MyTest{
// 	t := MyTest{myout,mya}
// 	return &t
// }


func main(){
	// c := make([]interface{},2)
	// c[0] = "send"
	// c[1] = "from"
	// out1 := COut(c,"cccc")
	// fmt.Println(out1.Out[0],out1.Out[1])
	// info,err := bitcoin.ValidateAddress("25EeN6fSpo8MrcUjERQmvpjdQfGkn8yND2")
	// if err == nil{
	// 	fmt.Println(info.ScriptPubKey)
	// }
	// info,err := bitcoin.GetTransaction("2e5753f438bde120eb01a7cf7656c3d055e77b30eb710e2cd11bfe9a7132750c")
	// if err==nil{
	// 	for i:=range info.Details {
	// 		fmt.Println(info.Details[i].Amount)
	// 	}
	// 	fmt.Println(info.Fee)
	// }
	// c := make(chan string,2)
	// c <- "helloworld"
	// c <- "maymay"
	// // c <- -199
	// close(c)
	// fmt.Printf("%s\n",<-c)
	// fmt.Printf("%d\n",<-c)
	// fmt.Printf("%d\n",<-c)
	// fmt.Printf("%d\n",<-c)

	// c <- 1

	// rtInfo := new (bitcoin.RawTransactionInfo)
	// rtInfo.TransactionInfo = &MyTransactionInfo{Txid:"3e05b2204b86b67618f2143cd9295106b69957614a4c4a30e51cd896651c7ffa",
	// 							Vout:425}
	// rtInfo.AmountInfo = {"3DzSVk4veMCkNbNT9CdETeE26uWxmNbBnD":0.00000888}

	// result,err := bitcoin.ValidateAddress("1P9U3cDzmuR5duJToaWbomyr2ckhvF4tqT")
	// if err==nil
	// {
	// 	fmt.Println(result)
	// }
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
	router.Run(":8080")

	uInfo := new(bitcoin.UnspentInfo)
	uInfo.Minconf = 0
	uInfo.Maxconf = 10
	uInfo.Address = []string{"3DzSVk4veMCkNbNT9CdETeE26uWxmNbBnD"}
	usInfo,err := bitcoin.ListUnspent(uInfo)
	if err == nil{
		fmt.Println(usInfo[0].Amount)
	}

	sbinfo,err := bitcoin.ListSinceBlock("","")
	if err == nil{
		fmt.Println(sbinfo.TranInfo[0].Address)
	}
	miningInfo,err := bitcoin.GetMiningInfo()
	if err == nil{
		fmt.Println(miningInfo.Networkhashps)
	}
	bestBlockHash,err := bitcoin.GetBestBlockHash()
	if  err==nil{
		fmt.Println(bestBlockHash)
	}
	accountAddress,err := bitcoin.GetAccountAddress("")
	if err ==nil{
		fmt.Println(accountAddress)
	}

	connectCount := bitcoin.GetConnectionCount()
	fmt.Println(connectCount)

	bInfo,err := bitcoin.GetBlock(bestBlockHash)
	if err==nil{
		fmt.Printf("%d\n",bInfo.Height)
	}
	blockHash,err := bitcoin.GetBlocHash(bInfo.Height)
	if err==nil{
		fmt.Println(blockHash)
	}
	status,err := ethereum.GetSyning()
	if (status && err==nil){
		datas,err := ethereum.GetHaveBalanceWithAddress()
		if err!= nil {
			fmt.Println(err)
			return
		}
		for i := range datas {
			fmt.Println("address:",datas[i].Address)
			fmt.Println("balance:",datas[i].Balance)
		}
	}else{
		amount ,err := ethereum.GetBalanceAmount()
		if err == nil{
			fmt.Println(amount)
		}

		blockNumber,err := ethereum.GetBlockNumber()
		if err==nil{
			fmt.Println("blockNumber",common.HexDec(blockNumber))
		}

		balance,err := ethereum.GetBalance("0x83a2533a81ee4ee55e219b0fab5016e723d12a42")
		if err == nil {
			fmt.Println("balance:",balance)
		}

		estimateGas,err := ethereum.GetEstimateGas("0x83a2533a81ee4ee55e219b0fab5016e723d12a42","0x0d5e7f601ee93b15b52288ed793da494cd759d30")
		if err == nil{
			fmt.Println("estimateGas",common.HexDec(estimateGas))
		}
	}
	
	// gasPrice,err := ethereum.GetGasPrice()
	// if err!= nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(common.HexDec(gasPrice))

	// blockNumber,err := ethereum.GetBlockNumber()
	// if err!= nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(common.HexDec(blockNumber))

}