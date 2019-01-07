package common 

import (
	"strconv"
)

//十六进制转十进制
//
func HexToDecimal(hex string)int64{
	if hex[:2]=="0x"{
		hex = hex[2:] 
	}
	result,err := strconv.ParseInt(hex,16,64)
	if err != nil{
		return 0
	}
	return result
}