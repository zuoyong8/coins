
package main


import (
	"fmt"
	"coins/ethereum"
	// "coins/bitcoin"
	// "reflect"
	//  "math"
	 "coins/common"
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
	c := make(chan string,2)
	c <- "helloworld"
	c <- "maymay"
	// c <- -199
	close(c)
	fmt.Printf("%s\n",<-c)
	// fmt.Printf("%d\n",<-c)
	// fmt.Printf("%d\n",<-c)
	// fmt.Printf("%d\n",<-c)

	// c <- 1

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