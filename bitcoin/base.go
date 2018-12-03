package bitcoin

import (
	"errors"
	"../rpc"
	"../config"
)

type CallFunc struct{
	method 	 string
	params	 []string
}

//
func New(method string,params []string) *CallFunc{
	api := CallFunc{method,params}
	return &api
}


//
func (cf *CallFunc) GetJosnBytes()([]byte){
	btcRpc,err := config.GetRpcInfo("btc")
	if err!=nil {
		return nil
	}

	client := rpc.New(btcRpc.Ip,btcRpc.Port,btcRpc.Username,btcRpc.Password)
	bytes,err := client.MakeRequest(cf.method,cf.params)
	if err!=nil {
		return nil
	}
	return bytes
}


//
func (cf *CallFunc) GetRpcBytes()([]byte,error){
	bytes := cf.GetJosnBytes()
	if bytes != nil {
		result,err := rpc.RpcJsonParse(bytes)
		if err !=nil{
			return nil,err
		}
		return result,nil
	}
	return nil,errors.New("not get data")
}