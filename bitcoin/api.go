package bitcoin

import (
	"encoding/json"
 	 "bytes"
)

//返回最近发生的与钱包有关的交易清单
func ListTransactions()([]TransInfo,error){
	callFunc := New("listtransactions",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return nil,err
	}
	var ts []TransInfo 
	err1 := json.Unmarshal(myBytes,&ts)
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
	myBytes,err := callFunc.GetRpcBytes()
	var addressInfo AddressInfo
	if err != nil{
		return addressInfo,err
	}
	err1 := json.Unmarshal(myBytes,&addressInfo)
	if err1 != nil {
		return addressInfo,err1
	}
	return addressInfo,nil
}

//返回具有帐户名称作为键，帐户余额作为值的对象
func ListAccounts()(map[string]interface{},error){
	callFunc := New("listaccounts",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil {
		return nil,err
	}
	var datas map[string]interface{}
	err1 := json.Unmarshal(myBytes,&datas)
	if err1 !=nil{
		return  nil,err1
	}
	return datas,nil
}

//版本0.7返回钱包中的所有地址和用于coincontrol的信息。
func ListAddressGroupings()([]interface{},error){
	callFunc := New("listaddressgroupings",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil {
		return nil,err
	}
	var datas []interface{}
	err1 := json.Unmarshal(myBytes,&datas)
	if err1 !=nil{
		return nil,err1
	}
	// for key,value := range datas{
	// }
	return datas,nil
}

//返回有关给定事务的对象
func GetTransaction(txid string)(TransactionInfo,error){
	params := make([]interface{},1)
	params[0] = txid
	callFunc := New("gettransaction",params)
	var info TransactionInfo 
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return info,nil
	}
	decoder  := json.NewDecoder(bytes.NewBuffer(myBytes))
	decoder.UseNumber()
	err = decoder.Decode(&info)
	// err = json.Unmarshal(myBytes,&info)
	if err != nil{
		return info,err
	}
	return info,nil
}

//用返回钱包的总体信息
func GetWalletInfo()(WalletInfo,error){
	callFunc := New("getwalletinfo",nil)
	myBytes,err := callFunc.GetRpcBytes()
	var w WalletInfo 
	if err != nil{
		return w,err
	}
	err1 := json.Unmarshal(myBytes,&w)
	if err1 != nil {
		return w,err1
	}
	return w,nil
}

//返回用于接收付款的新比特币地址
func GetNewaAddress()(string,error){
	callFunc := New("getnewaddress",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var newAddress string
	err1 := json.Unmarshal(myBytes,&newAddress)
	if err1 != nil {
		return "",err1
	}
	return newAddress,nil
}

//返回一个新的比特币地址，用于接收更改。这适用于原始交易，而非正常使用
func GetRawChangeAddress(account string)(string,error){
	callFunc := New("getrawchangeaddress",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err!= nil{
		return  "",err
	}
	var address string
	err1 := json.Unmarshal(myBytes,&address)
	if err1 != nil{
		return "",err1
	}
	return address,nil
}

//返回当前节点的总可用余额
func GetBalance()(float64,error){
	callFunc := New("getbalance",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0.00,err
	}
	var balance float64
	err1 := json.Unmarshal(myBytes,&balance)
	if err1 != nil {
		return 0.00,err1
	}
	return balance,nil
}


//根据哈哈希返回块相关信息
func GetBlock(hash string){
	params := make([]interface{},1)
	params[0] = hash
	callFunc := New("getblock",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{

	}
	if myBytes!=nil{}
}

//
func GetBlocHash(index int32){
	params := make([]interface{},1)
	params[0] = index
	callFunc := New("getblockhash",params)
	myBytes,err := callFunc.GetRpcBytes()
	if (err!=nil){

	}
	if myBytes!=nil{}
}


//返回最长块链中的块数。
func GetBlockCount()(int64,error){
	callFunc := New("getblockcount",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0,err
	}
	var blockCount int64
	err1 := json.Unmarshal(myBytes,&blockCount)
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
	myBytes,err := callFunc.GetRpcBytes()
	if err!= nil{
		return "",err
	}
	var txid string
	err = json.Unmarshal(myBytes,&txid)
	if err != nil {
		return "",err
	}
	return txid,nil
}

//
func SendToAddress(info SendInfo)(string,error){
	params := make([]interface{},2)
	params[0] = info.ToBitcoinAddress
	params[1] = info.Amount
	callFunc := New("sendtoaddress",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var txid string
	err1 := json.Unmarshal(myBytes,&txid)
	if err1 != nil {
		return "",err1
	}
	return txid,nil
}	


//从当前钱包转账至多个地址
func SendMany(smInfo SendManyInfo){
	params := make([]interface{},2)
	params[0] = smInfo.FromAccount
	params[1] = smInfo.SendsInfo

	// callFunc := New("sendmany",params)
}


//从一个帐户移动另一个钱包帐户
func Move(info MoveInfo)(bool,error){
	params := make([]interface{},3)
	params[0] = info.FromAccount
	params[1] = info.ToAccount
	params[2] = info.Amount

	callFunc := New("move",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return false,err
	}
	var status bool
	err1 := json.Unmarshal(myBytes,&status)
	if err1!=nil {
		return false,err1
	}
	return status,nil
}