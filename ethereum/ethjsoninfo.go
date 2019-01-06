package ethereum 

type GasInfo struct{
	Id 		int32 		`json:"id"`
	Jsonrpc string		`json:"jsonrpc"`
	Result	string 		`json:"result"`
}

type BlockByHashInfo struct{
	Number		string 		`json:"number"`
	Hash		string 		`json:"hash"`
	ParentHash	string 		`json:"parentHash"`
	Nonce		string 		`json:"nonce"`
	Sha3Uncles	string 		`json:"sha3Uncles"`
	LogsBloom	string 		`json:"logsBloom"`
	TransactionsRoot string  `json:"transactionsRoot"`
	StateRoot	string 		`json:"stateRoot"`
	Miner		string 		`json:"miner"`
	Difficulty  string 		`json:"difficulty"`
	TotalDifficulty string	`json:"totalDifficulty"`
	ExtraData	string 		`json:"extraData"`
	Size		string 		`json:"size"`
	GasLimit    string 		`json:"gasLimit"`
	GasUsed		string 		`json:"gasUsed"`
	Timestamp	string 		`json:"timestamp"`
	Transactions []interface{} `json:"transactions"`
	Uncles		[]string	`json:"uncles"`
}

type TransactionByHashInfo struct {
	Hash 				string 		`json:"hash"`
	Nonce				string 		 `json:"nonce"`
	BlockHash 			string 		`json:"blockHash"`
	BlockNumber 		string 		`json:"blockNumber"`
	TransactionIndex	string 		`json:"transactionIndex"`
	From 				string 		`json:"from"`
	To 					string 		`json:"to"`
	Value 				string 		`json:"value"`
	Gas 				string 		`json:"gas"`
	GasPrice			string 		`json:"gasPrice"`
	Input				string 		`json:"input"`
}

type FilterChangeInfo struct{
	LogIndex			string		`json:"logIndex"`
	BlockNumber			string		`json:"blockNumber"`
	BlockHash			string		`json:"blockHash"`
	TransactionHash		string		`json:"transactionHash"`
	TransactionIndex	string		`json:"transactionIndex"`
	Address				string		`json:"address"`
	Data				string		`json:"data"`
	Topics				[]string	`json:"topics"`
}

type TransactionReceiptInfo struct{
	TransactionHash		string 		`json:"transactionHash"`
	TransactionIndex	string		`json:"transactionIndex"`
	BlockNumber			string		`json:"blockNumber"`
	BlockHash			string		`json:"blockHash"`
	CumulativeGasUsed   string		`json:"cumulativeGasUsed"`
	GasUsed				string 		`json:"gasUsed"`
	ContractAddress		string		`json:"contractAddress"`
	Logs				[]interface{}	`json:"logs"`
	LogsBloom			string		`json:"logsBloom"`
	Status				string 		`json:"status"`
}

type WalletInfo struct{
	Url 				string 		`json:"url"`
	Status				string 		`json:"status"`
	Accounts			map[string]interface{} `json:"accounts"`
}