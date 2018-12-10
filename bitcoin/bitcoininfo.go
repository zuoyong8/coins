package bitcoin

type SendInfo struct{
	FromAccount 		string
	ToBitcoinAddress 	string
	Amount				float64
	Minconf				int32
	Comment				string
	Comment_To			string
}

type SendManyInfo struct {
	FromAccount string
	SendsInfo 	map[string]interface{}
	Minconf		int32
	Comment		string
}

type MoveInfo struct {
	FromAccount string
	ToAccount   string
	Amount		float64
	Minconf		int64
	Comment		string
}

type UnspentInfo struct{
	Minconf		int32
	Maxconf 	int32
	Address     []string
}

type MyTransactionInfo struct {
	Txid		string
	Vout		int32
}

type RawTransactionInfo struct {
	TransactionInfo []MyTransactionInfo
	AmountInfo 		map[string]interface{}
}
