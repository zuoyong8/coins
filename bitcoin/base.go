package bitcoin

import (
	"errors"
	"strings"

	"github.com/zuoyong8/coins/rpc"
	"github.com/zuoyong8/coins/config"
)

type CallFunc struct{
	method 	 string
	params	 []interface{}
}

//
func New(method string,params []interface{}) *CallFunc{
	api := CallFunc{method,params}
	return &api
}

//
func (cf *CallFunc) GetJosnBytes()([]byte){
	coinName := cf.method[:3]
	if (strings.Compare(coinName,"omn")==0){
		coinName = "usdt"
	}else{
		coinName = "btc"
	}
	btcRpc,err := config.GetCoinRpc(coinName)
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