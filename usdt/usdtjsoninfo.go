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