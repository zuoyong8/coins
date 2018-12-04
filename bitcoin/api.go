package bitcoin

import (
	"encoding/json"
// 	"reflect"
)

//返回最近发生的与钱包有关的交易清单
func ListTransactions()([]TransInfo,error){
	callFunc := New("listtransactions",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var ts []TransInfo 
	err1 := json.Unmarshal(bytes,&ts)
	if err1 != nil {
		return nil,err1
	}
	return ts,nil
}


//返回有关比特币地址的信息。
func ValidateAddress(address string)(AddressInfo,error){
	params := make([]interface{},1)
	params[0] = address
	callFunc := New("validateaddress",params)
	bytes,err := callFunc.GetRpcBytes()
	var addressInfo AddressInfo
	if err != nil{
		return addressInfo,err
	}
	err1 := json.Unmarshal(bytes,&addressInfo)
	if err1 != nil {
		return addressInfo,err1
	}
	return addressInfo,nil
}


//返回具有帐户名称作为键，帐户余额作为值的对象
func ListAccounts()(map[string]interface{},error){
	callFunc := New("listaccounts",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil {
		return nil,err
	}
	var datas map[string]interface{}
	err1 := json.Unmarshal(bytes,&datas)
	if err1 !=nil{
		return  nil,err1
	}
	return datas,nil
}


//版本0.7返回钱包中的所有地址和用于coincontrol的信息。
func ListAddressGroupings()([]interface{},error){
	callFunc := New("listaddressgroupings",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil {
		return nil,err
	}
	var datas []interface{}
	err1 := json.Unmarshal(bytes,&datas)
	if err1 !=nil{
		return nil,err1
	}
	// for key,value := range datas{
	// }
	return datas,nil
}


//用返回钱包的总体信息
func GetWalletInfo()(WalletInfo,error){
	callFunc := New("getwalletinfo",nil)
	bytes,err := callFunc.GetRpcBytes()
	var w WalletInfo 
	if err != nil{
		return w,err
	}
	err1 := json.Unmarshal(bytes,&w)
	if err1 != nil {
		return w,err1
	}
	return w,nil
}

//返回用于接收付款的新比特币地址
func GetNewaAddress()(string,error){
	callFunc := New("getnewaddress",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var newAddress string
	err1 := json.Unmarshal(bytes,&newAddress)
	if err1 != nil {
		return "",err1
	}
	return newAddress,nil
}

//返回当前节点的总可用余额
func GetBalance()(float64,error){
	callFunc := New("getbalance",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0.00,err
	}
	var balance float64
	err1 := json.Unmarshal(bytes,&balance)
	if err1 != nil {
		return 0.00,err1
	}
	return balance,nil
}


//返回最长块链中的块数。
func GetBlockCount()(int64,error){
	callFunc := New("getblockcount",nil)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0,err
	}
	var blockCount int64
	err1 := json.Unmarshal(bytes,&blockCount)
	if err1 != nil {
		return 0,err1
	}
	return blockCount,nil
}


//发送
func SendFrom(info SendInfo)(string,error){
	params := make([]interface{},3)
	params[0] = info.FromAccount
	params[1] = info.ToBitcoinAddress
	params[2] = info.Amount
	callFunc :=	New("sendfrom",params)
	bytes,err := callFunc.GetRpcBytes()
	if err!= nil{
		return "",err
	}
	var txid string
	err1 := json.Unmarshal(bytes,&txid)
	if err1 != nil {
		return "",err1
	}
	return txid,nil
}

//
func SendToAddress(info SendInfo)(string,error){
	params := make([]interface{},2)
	params[0] = info.ToBitcoinAddress
	params[1] = info.Amount
	callFunc := New("sendtoaddress",params)
	bytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var txid string
	err1 := json.Unmarshal(bytes,txid)
	if err1 != nil {
		return "",err1
	}
	return txid,nil
}	