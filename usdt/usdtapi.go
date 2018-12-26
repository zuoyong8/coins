package usdt

import (
	"github.com/zuoyong8/coins/bitcoin"
	
	"encoding/json"
	"strings"
)

//返回整个钱包的总令牌余额列表
func GetWalletBalances()([]WalletBalancesInfo,error){
	callFunc := bitcoin.New("omni_getwalletbalances",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var wbInfo []WalletBalancesInfo 
	err = json.Unmarshal(bytes,&wbInfo)
	if err != nil {
		return nil,err
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
	err = json.Unmarshal(bytes,&wabInfo)
	if err != nil {
		return nil,err
	}
	return wabInfo,nil
}

//获取有关Omni事务的详细信息
func GetTransaction(txid string)(TransactionInfo,error){
	params := make([]interface{},1)
	params[0] = txid
	callFunc := bitcoin.New("omni_gettransaction",params)
	bytes,err := callFunc.GetRpcBytes()
	var tranInfo TransactionInfo
	if err != nil{
		return tranInfo,err
	}
	err = json.Unmarshal(bytes,&tranInfo)
	if err != nil {
		return tranInfo,err
	}
	return tranInfo,nil
}


//列出钱包事务，可选地按地址和块边界过滤
func ListTransactions(address string)([]TransactionInfo,error){
	var params []interface{}
	if (strings.Compare(address,"all")==0){
		params = nil
	}else{
		params = make([]interface{},1)
		params[0] = address
	}
	callFunc := bitcoin.New("omni_listtransactions",params)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var tInfo []TransactionInfo 
	err = json.Unmarshal(bytes,&tInfo)
	if err != nil {
		return nil,err
	}
	return tInfo,nil
}


//返回客户端和协议的各种状态信息。
func Getinfo()(GetNodeinfo,error){
	callFunc := bitcoin.New("omn​​i_getinfo",nil)
	bytes,err := callFunc.GetRpcBytes()
	var nodeinfo GetNodeinfo
	if err != nil{
		return nodeinfo,err
	}
	err = json.Unmarshal(bytes,&nodeinfo)
	if err != nil {
		return nodeinfo,err
	}
	return nodeinfo,nil
}

//列出所有令牌或智能属性。
func ListProperties()([]PropertiesInfo,error){
	callFunc := bitcoin.New("omn​​i_listproperties",nil)
	bytes,err := callFunc.GetRpcBytes()
	var pInfo []PropertiesInfo
	if err != nil{
		return nil,err
	}
	err = json.Unmarshal(bytes,&pInfo)
	if err != nil {
		return nil,err
	}
	return pInfo,nil
}

//列出块中的所有Omni事务
func ListBlockTransactions(index int)([]string,error){
	params := make([]interface{},1)
	params[0] = index
	callFunc := bitcoin.New("omn​​i_listblocktransactions",params)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var transactions []string
	err = json.Unmarshal(bytes,&transactions)
	if err != nil {
		return nil,err
	}
	return transactions,nil
}
