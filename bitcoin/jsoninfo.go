package bitcoin

type TransInfo struct{
	Account  			string 		`json:"account"`
	Address  			string 		`json:"address"`
	Category 			string 		`json:"category"`
	Amount   			float32  	`json:"amount"`
	Label	 			string  	`json:"label"`
	Vout 	 			int32		`json:"vout"`
    Confirmations 		int32		`json:"confirmations"`
	Blockhash			string		`json:"blockhash"`
	Blockindex			int32		`json:"blockindex"`
	Blocktime			int32		`json:"blocktime"`
	Txid				string 		`json:"txid"`
	Walletconflicts		[]string	`json:"walletconflicts"`
	Time				int32		`json:"time"`
	Timereceived		int32		`json:"timereceived"`
    Bip125_replaceable  string   	`json:"bip125-replaceable"`
}


type WalletInfo struct{
	Walletname			string 		`json:"walletname"`
	Walletversio		int32		`json:"walletversion"`
	Balance				float32		`json:"balance"`
	Unconfirmed_balance float32		`json:"unconfirmed_balance"`
	Immature_balance	float32		`json:"immature_balance"`
	Txcount				int32		`json:"txcount"`
	Keypoololdest		int32		`json:"keypoololdest"`
	Keypoolsize			int32		`json:"keypoolsize"`
	Keypoolsize_hd_internal int32	`json:"keypoolsize_hd_internal"`
	Paytxfee			float32		`json:"paytxfee"`
	Hdmasterkeyid		string		`json:"hdmasterkeyid"`
}


type AddressInfo struct{
	IsValid				bool	    `json:"isvalid"`  //地址是否有效
	Address				string 		`json:"address"`
	ScriptPubKey		string		`json:"scriptPubKey"`
	IsMine				bool		`json:"ismine"`   //地址是否为本节点地址
	IsWatchOnly			bool 		`json:"iswatchonly"`
	IsScript			bool		`json:"isscript"`
	IsWitness			bool 		`json:"iswitness"`
	Script				string 		`json:"script"`
	EmbeddedInfo		AddressEmbeddedInfo `json:"embedded"`
	Account				string		`json:"account"`
	TimeStamp			int64		`json:"timestamp"`
	Addresses			[]string	`json:"addresses"`
	Hdkeypath			string		`json:"hdkeypath"`
	Hdmasterkeyid		string 		`json:"hdmasterkeyid"`
}


type AddressEmbeddedInfo struct {
	IsScript		bool 		`json:"isscript"`
	IsWitness		bool		`json:"iswitness"`
	Witness_Version int32		`json:"witness_version"`
	Witness_Program string		`json:"witness_program"`
	PubKey			string		`json:"pubkey"`
	Address			string 		`json:"address"`
	ScriptPubKey	string		`json:"scriptPubKey"`
}


type TransactionInfo struct{
	Amount				float64 	`json:"amount"`
	Fee					float64		`json:"fee"`
	Confirmations		int32		`json:"confirmations"`
	Blockhash			string 		`json:"blockhash"`
	Blockindex			int32		`json:"blockindex"`
	Blocktime			int32		`json:"blocktime"`
	Txid				string 		`json:"txid"`
	Walletconflicts		[]interface{} `json:"walletconflicts"`
	Time				int32		`json:"time"`
	TimeReceived		int32		`json:"timereceived"`
	Bip125_replaceable  string		`json:"bip125-replaceable"`
	Details		[]TransactionDetailsInfo	`json:"details"`
	Hex					string 		`json:"hex"`
}


type TransactionDetailsInfo struct{
	InvolvesWatchonly	bool 		`json:"involvesWatchonly"`
	Account				string 		`json:"account"`
	Address				string 		`json:"address"`
	Category			string		`json:"category"`
	Amount				float64		`json:"amount"`
	Label				string		`json:"label"`
	Vout				int32		`json:"vout"`
	Fee					float64		`json:"fee"`
	Abandoned			bool		`json:"abandoned"`
}


type BlockInfo struct {
	Hash				string 			`json:"hash"`
	Confirmations		int64			`json:"confirmations"`
	Strippedsize		int64			`json:"strippedsize"`
	Size				int64			`json:"size"`		
	Weight				int64			`json:"weight"`
	Height				int64			`json:"height"`
	Version				int64			`json:"version"`
	VersionHex			string			`json:"versionHex"`
	Merkleroot			string			`json:"merkleroot"`
	Tx					[]string		`json:"tx"`
	Time				int64			`json:"time"`
	Mediantime			int64			`json:"mediantime"`
	Nonce				int64			`json:"nonce"`
	Bits				string			`json:"bits"`
	Difficulty			float64			`json:"difficulty"`
	Chainwork			string			`json:"chainwork"`
	NTx					int64			`json:"nTx"`
	Previousblockhash	string			`json:"previousblockhash"`
	Nextblockhash		string			`json:"nextblockhash"`
}


type MiningInfo struct{
	Blocks				int64			`json:"blocks"`
	CurrentBlockWeight	int64			`json:"currentblockweight"`
	CurrentBlockTx		int64			`json:"currentblocktx"`
	Difficulty			float64			`json:"difficulty"`
	Networkhashps		interface{}		`json:"networkhashps"`
	Pooledtx			int64			`json:"pooledtx"`
	Chain				string			`json:"chain"`
	Warnings			string 			`json:"warnings"`
}


type SinceBlockInfo struct{
	TranInfo 			[]TransInfo			`json:"transactions"`
	Removed				[]interface{}		`json:"removed"`
	LastBlock			string 				`json:"lastblock"`
}


type UnSpentInfo struct{
	Txid				string				`json:"txid"`
	Vout				int32				`json:"vout"`
	Address				string				`json:"address"`
	Account				string				`json:"account"`
	ScriptPubKey		string				`json:"scriptPubKey"`
	Amount				float64				`json:"amount"`
	Confirmations		int32				`json:"confirmations"`
	Spendable			bool				`json:"spendable"`
	Solvable			bool				`json:"solvable"`
	Safe				bool				`json:"safe"`
}
