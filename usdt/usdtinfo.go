package usdt

type SendInfo struct{
	FromAddress 			string  //要发送的地址
	ToAddress				string  //接收者的地址
	PropertyId				int32   //要发送的令牌的标识符
	Amount					string  //发送金额
	RedeemAddress			string  //可以花费交易粉尘的地址
	ReferenceAmount			string  //发送到接收方的比特币金额
}

type SendsToInfo struct{
	FromAddress				string //要发送的地址
	PropertyId				int32  //要分发的令牌的标识符
	Amount					string//分配金额
	RedeemAddress			string//可以花费交易粉尘的地址（默认情况下为发件人
	DistributionProperty 	int32 //要分配给的财产持有人的标识符
}

type SendAllInfo struct{
	FromAddress				string//要发送的地址
	ToAddress				string//接收者的地址
	EcoSystem				int32 //要发送的令牌生态系统（1对于主要生态系统，2用于测试生态系统）
	RedeemAddress			string //可以花费交易粉尘的地址（默认情况下为发件人）
	ReferenceAmount			string //发送到接收方的比特币金额（默认为最小）
}

type FundedSendInfo struct{
	FromAddress				string//要发送的地址
	ToAddress				string//接收者的地址
	PropertyId				int32 //要发送的令牌的标识符
	Amount					string //发送金额
	FeeAddress				string//如果需要，用于支付费用的地址
}

type FundedSendallInfo struct {
	FromAddress				string //要发送的地址
	ToAddress				string //接收者的地址
	EcoSystem				int32  //要发送的令牌生态系统（主要生态系统为1，测试生态系统为2）
	FeeAddress				string //如果需要，用于支付费用的地址
}