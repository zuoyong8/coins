package config

import (
	"path/filepath"
	"os"
	"strings"
	"encoding/json"
	// "bytes"
	"errors"
	io "io/ioutil"
)

type RpcJsonInfo struct{
	Currency 	string		`json:"currency"`
	Ip 			string 		`json:"ip"`
	Port		int 		`json:"port"`
	Username	string		`json:"username"`
	Password 	string		`json:"password"`
}


//获取当前程序运行路径
func GetCurrRunDir() string{
	dir,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	} 
	return strings.Replace(dir,"\\","/",-1)
}

//
func GetCurrPath() string{
	return "D:/Codes/go/src/coins/config/config.json"
}


//根据币种从json配置文件中获取相应rpc信息
func GetRpcInfo(currency string) (*RpcJsonInfo, error){
	// jsonfile := GetCurrRunDir()+"/config/config.json"
	jsonfile := GetCurrPath()
	filedata,err := io.ReadFile(jsonfile)
	if err != nil{
		return nil,err
	}
	var jsondata []RpcJsonInfo
	err1 := json.Unmarshal(filedata,&jsondata)
	// decoder  := json.NewDecoder(bytes.NewBuffer(filedata))
	// err1 := decoder.Decode(&jsondata)
	if err1 != nil {
		return nil,err1
	}
	for i := range jsondata{
		if jsondata[i].Currency == currency {
			return &jsondata[i],nil
		}
	}
	return nil,errors.New("no the rpcinfo")
}