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