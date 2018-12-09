package usdt

import (
	"encoding/json"
	"coins/bitcoin"
)

//返回整个钱包的总令牌余额列表
func GetWalletBalances()([]WalletBalancesInfo,error){
	callFunc := bitcoin.New("omni_getwalletbalances",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var wbInfo []WalletBalancesInfo 
	err1 := json.Unmarshal(bytes,&wbInfo)
	if err1 != nil {
		return nil,err1
	}
	return wbInfo,nil
}

//返回每个钱包地址的所有令牌余额列表
func GetWalletaddressBalances()([]WalletaddressBalancesInfo,error){
	callFunc := bitcoin.New("omni_getwalletaddressbalances",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var wabInfo []WalletaddressBalancesInfo 
	err1 := json.Unmarshal(bytes,&wabInfo)
	if err1 != nil {
		return nil,err1
	}
	return wabInfo,nil
}
