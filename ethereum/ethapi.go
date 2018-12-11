package ethereum

import (
	"../common"
)

//获取当前钱包内所有地址
func GetAccouts()([]string,error){
	callFunc,err := New("eth_accounts",nil)
	if err != nil{
		return nil,err
	}
	var result []string
	err = callFunc.EthClient.Call(&result,callFunc.Method)
	if err!=nil{
		return nil,err
	}
	return result,nil
}


//获取当前节点eth总数量
func GetBalanceAmount()(common.Decimal,error){
	result,err :=  GetHaveBalanceWithAddress()
	amount := common.New(0.00,0)
	if err!=nil {
		return amount,err
	}
	for i:=range result{
		amount = amount.Add(result[i].Balance)
	}
	return amount,nil
}


//获取钱包有eth的地址
func GetHaveBalanceWithAddress()([]BalanceInfo,error){
	accounts,err := GetAccouts()
	if err!= nil {
		return  nil,err
	}
	var bInfo []BalanceInfo
	for i := range accounts {
		b,err := GetBalance(accounts[i])
		if err==nil{
			balance := common.HexDec(b)
			if balance >0 {
				d := common.New(balance,Ether)
				ethBalance := d.Rescale(Ether)
				info := BalanceInfo{accounts[i],ethBalance}
				bInfo = append(bInfo,info)
			}
		}
	}
	return bInfo,nil
}


//从钱包里地址获取eth数量
func GetBalance(address string)(string,error){
	Params := make([]interface{},2)
	Params[0] = address
	Params[1] = "latest"
	callFunc,err := New("eth_getBalance",Params)
	if err != nil {
		return "",err
	}
	
	var balance string
	err = callFunc.EthClient.Call(&balance,callFunc.Method,Params[0],Params[1])
	if err != nil{
		return "",err
	}
	return balance,nil
}

//获取gas价格
func GetGasPrice()(string,error){
	callFunc,err := New("eth_gasPrice",nil)
	if err != nil{
		return "",err
	}
	var gasPrice string
	err = callFunc.EthClient.Call(&gasPrice,callFunc.Method)
	if err!=nil{
		return "",err
	}
	return gasPrice,nil
}


//生成并返回允许事务完成所需的气体估计值
func GetEstimateGas(fromAddress string,toAddress string)(string,error){
	Params := make(map[string]interface{},2)
	Params["from"] = fromAddress
	Params["to"] = toAddress
	callFunc,err := New("eth_estimateGas",nil)
	if err != nil {
		return "",err
	}
	var estimateGas string
	err = callFunc.EthClient.Call(&estimateGas,callFunc.Method,Params)
	if err!=nil{
		return "",err
	}
	return estimateGas,nil
}


//获取区块同步状态
func GetSyning()(bool,error){
	callFunc,err := New("eth_syncing",nil)
	if err != nil{
		return false,err
	}
	var status bool
	err = callFunc.EthClient.Call(&status,callFunc.Method)
	if err!=nil{
		return false,err
	}
	return status,nil
}


//获取当前节点同步的最新区块高度
func GetBlockNumber()(string,error){
	callFunc,err := New("eth_blockNumber",nil)
	if err != nil{
		return "",err
	}
	var blockNumber string
	err = callFunc.EthClient.Call(&blockNumber,callFunc.Method)
	if err!=nil{
		return "",err
	}
	return blockNumber,nil
}

//返回指定地址发生的交易数量
func GetTransactionCount(data string)(string,error){
	Params := make([]interface{},2)
	Params[0] = data
	Params[1] = "latest"
	callFunc,err := New("eth_getTransactionCount",Params)
	if err != nil {
		return "",err
	}
	
	var result string
	err = callFunc.EthClient.Call(&result,callFunc.Method,Params[0],Params[1])
	if err != nil{
		return "",err
	}
	return result,nil
}

//返回具有指定哈希的块
func GetBlockByHash(dataHash string)(BlockByHashInfo,error){
	Params := make([]interface{},2)
	Params[0] = dataHash
	Params[1] = true
	callFunc,err := New("eth_getBlockByHash",Params)
	var info BlockByHashInfo
	if err != nil {
		return info,err
	}
	err = callFunc.EthClient.Call(&info,callFunc.Method,Params[0],Params[1])
	if err != nil{
		return info,err
	}
	return info,nil
}

//返回指定哈希对应的交易
func GetTransactionByHash(dataHash string)(TransactionByHashInfo,error){
	Params := make([]interface{},1)
	Params[0] = dataHash
	callFunc,err := New("eth_getTransactionByHash",Params)
	var tranInfo TransactionByHashInfo
	if err != nil {
		return tranInfo,err
	}
	err = callFunc.EthClient.Call(&tranInfo,callFunc.Method,Params[0])
	if err != nil{
		return tranInfo,err
	}
	return tranInfo,nil
}
//发送
// func SendTransaction(fromAddress string, toAddress string, ether float64 )(txid string ,error){
// 	Params 


// }

