package config

import (
	"path/filepath"
	"os"
	"strings"
	"encoding/json"
	"bytes"
	"fmt"
	io "io/ioutil"
)

//获取当前程序运行路径
func GetCurrRunDir() string{
	dir,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir,"\\","/",-1)
}


//根据币种从json配置文件中获取相应rpc信息
func GetRpcInfo(currency string) (map[string]interface{}, error){
	jsonfile := GetCurrRunDir()+"/config/config.json"
	filedata,err := io.ReadFile(jsonfile)
	if err != nil{
		return nil,err
	}
	var jsondata map[string]interface{}
	decoder  := json.NewDecoder(bytes.NewBuffer(filedata))
	err1 := decoder.Decode(&jsondata)
	if err1 != nil {
		return nil,err1
	}
	data := jsondata["rpcserver"].(map[string]interface{})
	if data != nil{
		fmt.Println(data)
	}
	return data[currency].(map[string]interface{}),nil
}