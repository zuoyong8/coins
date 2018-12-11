package usdt

type WalletBalancesInfo struct {
	Propertyid		int32		`json:"propertyid"`
	Name			string		`json:"name"`
	Balance			string		`json:"balance"`
	Reserved		string 		`json:"reserved"`
	Frozen			string		`json:"frozen"`
}

type WalletaddressBalancesInfo struct{
	Address			string 		`json:"address"`
	Balances		[]WalletBalancesInfo `json:"balances"`
}

type TransactionInfo struct{
	TxId   				string		`json:"txid"`
	Fee			    	string		`json:"fee"`
	SendingAddress		string		`json:"sendingaddress"`
	ReferenceAddress	string		`json:"referenceaddress"`
	IsMine				bool		`json:"ismine"`
	Version				int32		`json:"version"`
	Type_Int			int32		`json:"type_int"`
	Type				string		`json:"type"`
	PropertyId			int32		`json:"propertyid"`
	Divisible			bool		`json:"divisible"`
	Amount				string		`json:"amount"`
	Valid				bool		`json:"valid"`
	BlockHash			string		`json:"blockhash"`
	BlockTime			int64		`json:"blocktime"`
	PositioninBlock		int32		`json:"positioninblock"`
	Block				int64		`json:"block"`
	Confirmations		int32		`json:"confirmations"`
}

type GetNodeinfo struct{
	Omnicoreversion_int	int64 		`json:"omnicoreversion_int"`  
	Omnicoreversion		string		`json:"omnicoreversion"`
	Mastercoreversion	string 		`json:"mastercoreversion"`
	Bitcoincoreversion	string		`json:"bitcoincoreversion"`
	Block				int64		`json:"block"`
	Blocktime			int64		`json:"blocktime"`
	Blocktransactions	int32		`json:"blocktransactions"`
	Totaltrades			int32		`json:"totaltrades"`
	Totaltransactions	int64		`json:"totaltransactions"`
	Alerts				[]interface{} `json:"alerts"` 
}

type PropertiesInfo struct {
	Propertyid 			int64 		`json:"propertyid"`
	Name 				string		`json:"name"`
	Category			string		`json:"category"`
	Subcategory			string		`json:"subcategory"`
	Data 				string		`json:"data"`
	Url 				string		`json:"url"`
	Divisible 			bool 		`json:"divisible"`
}