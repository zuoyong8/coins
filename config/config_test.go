package config

import (
	"testing"
)

func TestConfig(t *testing.T){
	result ,err := GetCoinRpc("btc")
	if err!=nil{
		t.Error(err)
	}
	if result != nil{
		
	}
}