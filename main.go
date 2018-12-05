package main


import (
	"fmt"
	"./ethereum"
	"./common"
	// "./bitcoin"
)


func main(){
	// datas,err := ethereum.GetAccouts()
	// if err!= nil {
	// 	fmt.Println(err)
	// }
	// for i:=range datas {
	// 	// fmt.Println(datas[i])
	// }

	gasPrice,err := ethereum.GetGasPrice()
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(common.HexDec(gasPrice))

	blockNumber,err := ethereum.GetBlockNumber()
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(common.HexDec(blockNumber))

}