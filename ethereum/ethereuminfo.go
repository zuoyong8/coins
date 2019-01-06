package ethereum

import (
	"github.com/zuoyong8/coins/common"
)

type BalanceInfo struct{
	Address 	string
	Balance 	common.Decimal
}

type TransactionInfo struct{
	From 		string
	To			string
	Value		string
	Data		string
}