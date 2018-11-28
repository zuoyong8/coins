package main


import (
	"fmt"
	"./rpc"
)


func main(){
	c := rpc.NewClient("sjyy", "Sjyy2018", "192.168.0.185", 8334)
	if c == nil {
		fmt.Println("null")
		return
	}

	params := []string{}
	bytes, err := c.MakeRequest("getwalletinfo", params)
	if err != nil {
		fmt.Println(err)
		return 
	}
	result,error := c.JsonParseToMapString(bytes)
	if error != nil {
		fmt.Println(error)
		return 
	}
	
	for n, a := range result {
        fmt.Printf("%v: %v\n", n, a)
    }
	// fmt.Println(result)

}