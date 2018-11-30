package bitcoin

import (
	"../rpc"
	"../config"
)

type API struct{
	method 	 string
	params	 []string
}


func New(method string,params []string) *API{
	api := API{method,params}
	return &api
}

//
func (api *API) GetJosnBytes()([]byte){
	btcRpc,err := config.GetRpcInfo("btc")
	if err!=nil {
		return nil
	}

	client := rpc.New(btcRpc.Ip,btcRpc.Port,btcRpc.Username,btcRpc.Password)
	bytes,err := client.MakeRequest(api.method,api.params)
	if err!=nil {
		return nil
	}
	return bytes
}