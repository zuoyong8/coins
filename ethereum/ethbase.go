package ethereum

import (
	"github.com/ethereum/go-ethereum/rpc"
	"../config"
	myrpc "../rpc"
)

type CallFunc struct{
	Method		string
	Params 		[]interface{}
	EthClient	*rpc.Client
}

func New(method string,params []interface{})(*CallFunc,error){
	ethRpcInfo,err := config.GetRpcInfo("eth")
	if err != nil {
		return nil,err
	}
	client,err := myrpc.NewEthClient(ethRpcInfo.Ip,ethRpcInfo.Port)
	if err != nil {
		return nil,err
	}
	callFunc := CallFunc{method,params,client}
	return &callFunc,nil
}
