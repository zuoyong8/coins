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




