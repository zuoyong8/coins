package tests

import (
	"testing"
	"fmt"
	"../../bitcoin"
	// "reflect"
)

func TestListAccounts(t *testing.T){
	datas,err := bitcoin.ListAccounts()
	if err != nil{
		t.Error(err)
		return
	}
	for key,value := range datas {
		fmt.Println("key:",key)
		fmt.Println("value:",value)
	}
}


func TestValidateAddress(t *testing.T){
	datas,err := bitcoin.ValidateAddress("1A92GKYakBSETpnAfZG5sVbRsjwxr7k83y")
	if err!=nil {
		t.Error(err)
		return 
	}
	fmt.Println(datas.Address)
}

func TestGetBalance(t *testing.T){
	balance,err := bitcoin.GetBalance()
	if err != nil{

	}
	fmt.Println(balance)
}


func TestGetTransaction(t *testing.T){
	info,err := bitcoin.GetTransaction("2e5753f438bde120eb01a7cf7656c3d055e77b30eb710e2cd11bfe9a7132750c")
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(info.Hex)
	//fmt.Println(info.tranDetailsInfo.Amount)
}


func TestListAddressGroupings(t *testing.T){
	datas,err := bitcoin.ListAddressGroupings()
	if err != nil{
		t.Error(err)
		return
	}
	for key,value := range datas {
		fmt.Println("key:",key)
		// fmt.Println(reflect.TypeOf(value))
		// fmt.Println(reflect.ValueOf(value))
		// fmt.Println("value:",value)
		for key1,value1 := range value.([]interface{}){
			fmt.Println("key1:",key1)
			// fmt.Println("value1:",value1)
			for key2,value2 := range value1.([]interface{}){
				fmt.Println("key2:",key2)
				fmt.Println("value2:",value2)
			}
		}
	}
}