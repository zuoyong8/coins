package tests

import (
	"github.com/zuoyong8/coins"
	
	"testing"
	"fmt"
	// "../../common"
)

func TestCoins(t *testing.T){
	db,err := models.InitDB()
	if err ==nil {
		// coins :=  &models.Coins{
		// 	Currency:"eth",
		// 	FromAddress:"0xwiuieoewond",
		// 	ToAddress:"0xioeuqoiue134",
		// 	Amount:0.05,
		// 	TranferTime:common.GetCurrentTime(),
		// 	Confirmations:1024,
		// 	BlockNumber:20334,
		// 	Txid:"0x13333413"}

		// err = coins.Insert()
		// if err == nil{

		// }
		result,err1 := models.GetCoinsByCurrency("eth")
		if err1 == nil{
			for i:= range result{
				fmt.Println(result[i].Amount)
			}
		}
		// c,err1 := models.GetCoinsByTxid("eth")
		// if err1==nil{
		// 	fmt.Println(c.Amount)
		// }
	}
	defer db.Close()
}