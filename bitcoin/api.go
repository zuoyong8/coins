package bitcoin

import (
	"fmt"
)


// func GetWallerInfo
func ListTransactions()[]*TransInfo{
	api := New("listtransactions",nil)
	bytes := api.GetJosnBytes()

	if bytes != nil {
		fmt.Println(bytes)
	}

	return nil
	// result ,err := c.JsonParseToArray(bytes)
	// if err!=nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// return result
	// for _, v:= range result{
	// 	item := v.(map[string]interface{})
	// 	fmt.Println(item["confirmations"])
	// }
}