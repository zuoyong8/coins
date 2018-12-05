package ethereum


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

