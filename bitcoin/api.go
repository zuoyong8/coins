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
	err = json.Unmarshal(myBytes,&ts)
	if err != nil {
		return nil,err
	}
	return ts,nil
}


//获取所有事务
func ListSinceBlock(blockHash string,targetConfirmations string)(SinceBlockInfo,error){
	params := make([]interface{},2)
	if len(blockHash)> 0 || len(targetConfirmations) >0 {
		params[0] = blockHash
		params[1] = targetConfirmations
	}else{
		params=nil
	}
	callFunc := New("listsinceblock",params)
	myBytes,err := callFunc.GetRpcBytes()
	var sbInfo SinceBlockInfo
	if err != nil{
		return sbInfo,err
	}
	err = json.Unmarshal(myBytes,&sbInfo)
	if err != nil {
		return sbInfo,err
	}
	return sbInfo,nil
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
	err = json.Unmarshal(myBytes,&addressInfo)
	if err != nil {
		return addressInfo,err
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
	err = json.Unmarshal(myBytes,&datas)
	if err !=nil{
		return  nil,err
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
	err = json.Unmarshal(myBytes,&datas)
	if err !=nil{
		return nil,err
	}
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

//返回钱包的总体信息
func GetWalletInfo()(WalletInfo,error){
	callFunc := New("getwalletinfo",nil)
	myBytes,err := callFunc.GetRpcBytes()
	var w WalletInfo 
	if err != nil{
		return w,err
	}
	err = json.Unmarshal(myBytes,&w)
	if err != nil {
		return w,err
	}
	return w,nil
}

//获取与其他节点的连接数
func GetConnectionCount()int32{
	callFunc := New ("getconnectioncount",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0
	}
	var connectionCount int32
	err = json.Unmarshal(myBytes,&connectionCount)
	if err != nil {
		return 0
	}
	return connectionCount
}


//返回用于接收付款的新比特币地址
func GetNewaAddress()(string,error){
	callFunc := New("getnewaddress",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var newAddress string
	err = json.Unmarshal(myBytes,&newAddress)
	if err != nil {
		return "",err
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
	err = json.Unmarshal(myBytes,&address)
	if err != nil{
		return "",err
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
	err = json.Unmarshal(myBytes,&balance)
	if err != nil {
		return 0.00,err
	}
	return balance,nil
}


//根据块哈希获取块相关信息
func GetBlock(hash string)(BlockInfo,error){
	params := make([]interface{},1)
	params[0] = hash
	callFunc := New("getblock",params)
	myBytes,err := callFunc.GetRpcBytes()
	var bInfo BlockInfo
	if err != nil{
		return bInfo,err
	}
	err = json.Unmarshal(myBytes,&bInfo)
	if err != nil {
		return bInfo,err
	}
	return bInfo,nil
}

//返回包含与挖掘相关的信息
func GetMiningInfo()(MiningInfo,error){
	callFunc := New("getmininginfo",nil)
	myBytes,err := callFunc.GetRpcBytes()
	var mInfo MiningInfo
	if err != nil{
		return mInfo,err
	}
	err = json.Unmarshal(myBytes,&mInfo)
	if err != nil {
		return mInfo,err
	}
	return mInfo,nil
}

//根据块索引获取块哈希 
func GetBlocHash(index int64)(string,error){
	params := make([]interface{},1)
	params[0] = index
	callFunc := New("getblockhash",params)
	myBytes,err := callFunc.GetRpcBytes()
	if (err!=nil){
		return "",err
	}
	var blockHash string
	err = json.Unmarshal(myBytes,&blockHash)
	if err!=nil{
		return "",err
	}
	return blockHash,nil
}

//返回最长块链中最佳（tip）块的哈希值
func GetBestBlockHash()(string,error){
	callFunc := New("getbestblockhash",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var bestBlockHash string
	err = json.Unmarshal(myBytes,&bestBlockHash)
	if err != nil {
		return "",err
	}
	return bestBlockHash,nil
}


//返回最长块链中的块数。
func GetBlockCount()(int64,error){
	callFunc := New("getblockcount",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return 0,err
	}
	var blockCount int64
	err = json.Unmarshal(myBytes,&blockCount)
	if err != nil {
		return 0,err
	}
	return blockCount,nil
}


//返回用于接收此帐户付款的当前比特币地址。如果<account>不存在，它将与将返回的相关新地址一起创建
func GetAccountAddress(account string)(string,error){
	params := make([]interface{},1)
	params[0] = account
	callFunc := New("getaccountaddress",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var accountAddress string
	err = json.Unmarshal(myBytes,&accountAddress)
	if err != nil {
		return "",err
	}
	return accountAddress,nil
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
	err = json.Unmarshal(myBytes,&txid)
	if err != nil {
		return "",err
	}
	return txid,nil
}	


//从当前钱包转账至多个地址
func SendMany(smInfo *SendManyInfo){
	params := make([]interface{},2)
	params[0] = smInfo.FromAccount
	params[1] = smInfo.SendsInfo
	// callFunc := New("sendmany",params)
}


//从一个帐户移动另一个钱包帐户
func Move(info *MoveInfo)(bool,error){
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
	err = json.Unmarshal(myBytes,&status)
	if err!=nil {
		return false,err
	}
	return status,nil
}


//设置与给定地址关联的帐户。分配已分配给同一帐户的地址将创建与该帐户关联的新地址
func SetAccount(bitcoinAddress string,account string)string{
	params := make([]interface{},2)
	params[0] = bitcoinAddress
	params[1]=account
	callFunc := New("setaccount",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return ""
	}
	var newAddress string
	err = json.Unmarshal(myBytes,&newAddress)
	if err != nil {
		return ""
	}
	return newAddress
}

//停止比特币服务器
func Stop()string{
	callFunc := New("stop",nil)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return ""
	}
	var message string
	err = json.Unmarshal(myBytes,&message)
	if err != nil {
		return ""
	}
	return message
}


////////----UTXO-------//////////
/////////////////////////////////
//返回钱包中未使用的事务输出数组
func ListUnspent(info *UnspentInfo)([]UnSpentInfo,error){
	params := make([]interface{},3)
	params[0] = info.Minconf
	params[1] = info.Maxconf
	params[2] = info.Address
	callFunc := New("listunspent",params)
	myBytes,err := callFunc.GetRpcBytes()
	var usInfo []UnSpentInfo
	if err != nil{
		return nil,err
	}
	err = json.Unmarshal(myBytes,&usInfo)
	if err != nil {
		return nil,err
	}
	return usInfo,nil
}


//创建待发送交易
func CreateRawTransaction(info RawTransactionInfo)(string,error){
	params := make([]interface{},2)
	params[0] = info.TransactionInfo
	params[1] = info.AmountInfo

	callFunc := New("createrawtransaction",params)
	myBytes,err := callFunc.GetRpcBytes()
	if err != nil{
		return "",err
	}
	var hexEncoded string
	err = json.Unmarshal(myBytes,&hexEncoded)
	if err != nil {
		return "",err
	}
	return hexEncoded,nil
}