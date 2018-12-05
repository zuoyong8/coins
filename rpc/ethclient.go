package rpc

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewEthClient(host string,port int)(*rpc.Client,error){
	baseUrl := fmt.Sprintf("http://%s:%d", host, port)
	client,err := rpc.Dial(baseUrl)
	if err != nil {
		return nil,err
	}
	return client,err
}