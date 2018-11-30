package main


import (
	"fmt"
	// "./bitcoin"
	"./config"
	// "./rpc"
)


func main(){
	jsondata,err := config.GetRpcInfo("btc")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(jsondata["username"])
	// rpcclient := rpc.NewClient(jsondata["username"],jsondata["password"],jsondata["ip"],jsondata["port"])
	// fmt.Println(bitcoin.ListTransactions())
}