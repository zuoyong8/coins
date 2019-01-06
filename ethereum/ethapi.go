package ethereum

import (
	"github.com/zuoyong8/coins/common"
)

//geth
//https://ethfans.org/posts/ethereum-token-standards-list-1
//datadir  		--- 设置当前区块链网络数据存放的位置
//nodiscover	--- 私有链地址
//console		--- 启动命令行模式，可以在Geth中执行命令
//identity		--- 区块链的标示，用于标示目前网络的名字
//rpc			--- 开启rpc通道
//rpcapi 		--- 要开放哪些rpc api(db,eth,net,web3,personal)
	// eth：包含一些跟操作区块链相关的方法
	// net：包含以下查看p2p网络状态的方法
	// admin：包含一些与管理节点相关的方法
	// miner：包含启动&停止挖矿的一些方法
	// personal：主要包含一些管理账户的方法
	// txpool：包含一些查看交易内存池的方法
	// web3：包含了以上对象，还包含一些单位换算的方法
//rpccorsdomain --- 允许能连接到你的节点执行rpc api的url，使用逗号分隔。*表示任何url都可以连接
//rpcaddr		--- HTTP-RPC服务器接口地址，默认为localhost
//rpcport       --- HTTP-RPC服务器端口地址，默认为8545
//networkid		--- 网络标识，私有链取一个大于4的随意的值

//
func Coinbase()(string,error){
	callFunc,err := New("eth_coinbase",nil)
	if err != nil{
		return "",err
	}
	var result string
	err = callFunc.EthClient.Call(&result,callFunc.Method)
	if err!=nil{
		return "",err
	}
	return result,nil
}

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


//发送交易
//代币转账token_transfer
//value：由于是发送代币，这里为0
//data：附加的消息。这里由合约中transfer方法，方法参数一(接收方地址)，方法参数二(代币数量)的十六进制组成
//data值应该为方法名的sha3的前8个字节+参数的64字节，不够前面补充为0。
//web3.sha3("transfer(address,uint256)")
//"data": "0xa9059cbb000000000000000000000000696d69b81c6bdf6d46ddb66ee2175df7f9de7c4600000000000000000000000000000000000000000000000ad78ebc5ac6200000"
func SendTransaction(info TransactionInfo)(string,error){
	gasPrice,err := GetGasPrice()
	gasLimit,err1 := GetEstimateGas(info.From,info.To)
	if err!=nil{
		return "",err
	}
	if err!=nil{
		return "",err1
	}
	Params := make([]interface{},5)
	Params[0] = info.From
	Params[1] = info.To
	Params[2] = gasLimit
	Params[3] = gasPrice
	Params[4] = info.Value
	Params[5] = info.Data
	callFunc,err := New("eth_sendTransaction",Params)
	var txHtash string
	if err != nil {
		return "",err
	}
	err = callFunc.EthClient.Call(&txHtash,callFunc.Method,Params[0],Params[1],Params[2],Params[3],Params[4],Params[5])
	if err != nil{
		return "",err
	}
	return txHtash,nil
}

//获取指定地址代币余额--要调用的方法名balanceOf和指定地址的十六进制
//balanceOf(address _owner)
//获取代币小数位--要调用的方法名decimals的十六进制
// HEX String - 指定区块号的十六进制
// String "earliest" - 表示最初或创世区块号
// String "latest" - 表示最新挖出的区块号
// String "pending" - 表示pending状态的交易
func EthCall(to string,data string)(string,error){
	Params := make([]interface{},3)
	Params[0] = to
	Params[1] = data
	Params[2] = "latest"
	callFunc,err := New("eth_call",Params)
	var result string
	if err != nil {
		return "",err
	}
	err = callFunc.EthClient.Call(&result,callFunc.Method,Params[0],Params[1],Params[2])
	if err != nil{
		return "",err
	}
	return result,nil
}

//获取指定高度的区块详情
func GetBlockByNumber(dataHash string)(BlockByHashInfo,error){
	Params := make([]interface{},2)
	Params[0] = dataHash
	Params[1] = true
	callFunc,err := New("eth_getBlockByNumber",Params)
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

//开启挖矿
func MinerStart()error{
	callFunc,err := New("miner_start ",nil)
	if err != nil {
		return err
	}
	var result string
	err = callFunc.EthClient.Call(&result,callFunc.Method)
	if err != nil{
		return err
	}
	return nil
}

//停止挖矿
func MinerStop()error{
	callFunc,err := New("miner_stop ",nil)
	if err != nil {
		return err
	}
	var result string
	err = callFunc.EthClient.Call(&result,callFunc.Method)
	if err != nil{
		return err
	}
	return nil
}

////personal钱包方法
//创建账户
func PersonalNewAccount()(string,error){
	callFunc,err := New("personal_newAccount",nil)
	if err != nil {
		return "",err
	}
	var newAddress string
	err = callFunc.EthClient.Call(&newAddress,callFunc.Method)
	if err != nil{
		return "",err
	}
	return newAddress,nil
}

//解锁账号
func PersonalUnlockAccount(address string)(bool,error){
	Params := make([]interface{},3)
	Params[0] = address
	Params[1] = "PASSWORD"
	Params[2] = 30
	callFunc,err := New("personal_unlockAccount",Params)
	var isOk bool
	if err != nil {
		return false,err
	}
	err = callFunc.EthClient.Call(&isOk,callFunc.Method,Params[0],Params[1],Params[2])
	if err != nil{
		return false,err
	}
	return isOk,nil
}

//获取所有本地账户地址
func PersonalListAccounts()([]string,error){
	callFunc,err := New("personal_listAccounts ",nil)
	var addresses []string
	if err != nil {
		return nil,err
	}
	err = callFunc.EthClient.Call(&addresses,callFunc.Method)
	if err != nil{
		return nil,err
	}
	return addresses,nil
}

// 获取所有本地钱包信息
func PersonalListWallets()([]WalletInfo,error){
	callFunc,err := New("personal_listWallets  ",nil)
	var walletInfos []WalletInfo
	if err != nil {
		return nil,err
	}
	err = callFunc.EthClient.Call(&walletInfos,callFunc.Method)
	if err != nil{
		return nil,err
	}
	return walletInfos,nil
}

//通过私钥和密码导入keystore文件
//返回账户地址
func PersonalImportRawKey(privateKey string,password string)(string,error){
	Params := make([]interface{},2)
	Params[0] = privateKey
	Params[1] = password
	callFunc,err := New("personal_importRawKey",Params)
	var result string
	if err != nil {
		return "",err
	}
	err = callFunc.EthClient.Call(&result,callFunc.Method,Params[0],Params[1])
	if err != nil{
		return "",err
	}
	return result,nil
}

//在节点中创建一个过滤器，以便当新块生成时进行通知
func NewBlockFilter()(string,error){
	callFunc,err := New("eth_newBlockFilter",nil)
	if err != nil {
		return "",err
	}
	var filterCode string
	err = callFunc.EthClient.Call(&filterCode,callFunc.Method)
	if err != nil{
		return "",err
	}
	return filterCode,nil
}

//轮询指定的过滤器，并返回自上次轮询之后新生成的日志数组
//1、对于使用eth_newBlockFilter返回创建的过滤器是块哈希
//2、对于使用eth_newPendingTransactionFilter 返回创建的过滤器是事务哈希
//3、对于使用eth_newFilter日志创建的过滤器
func GetFilterChanges(filterType FilterType,filterCode string)([]interface{},error){
	Params := make([]interface{},1)
	Params[0] = filterCode
	callFunc,err := New("eth_getFilterChanges",Params)
	if err != nil {
		return nil,err
	}
	switch filterType {
		case BlockFilter:
			var result []interface{}
			err = callFunc.EthClient.Call(&result,callFunc.Method,Params[0])
			if err != nil{
				return  nil,err
			}
			return result,nil
	}
	return nil,nil
}

//发送
// func SendTransaction(fromAddress string, toAddress string, ether float64 )(txid string ,error){
// 	Params 
// }

