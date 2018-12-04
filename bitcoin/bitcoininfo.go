package bitcoin

type SendInfo struct{
	FromAccount 		string
	ToBitcoinAddress 	string
	Amount				float64
	Minconf				int32
	Comment				string
	Comment_To			string
}


type MoveInfo struct {
	FromAccount string
	ToAccount   string
	Amount		float64
	Minconf		int64
	Comment		string
}
// type SendManyInfo struct {
// 	FromAccount 		string
// 	AmountInfo			map[string]
// }