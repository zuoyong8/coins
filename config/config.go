package config

import (
	"path/filepath"
	"os"
	"strings"
	"encoding/json"
	"fmt"
	// "bytes"
	"errors"
	io "io/ioutil"
	"github.com/spf13/viper"
)

type RpcConnectInfo struct{
	Currency 	string				
	Ip 			string 			
	Port		int				
	Username	string			
	Password 	string			
}

type DbConnectInfo struct{
	Ip				string	
	Port			int			
	Username		string		
	Password		string		
	Dbname			string	
}


func SetViperConfig(){
    viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")
	viper.SetConfigName("coins")
}

//获取当前程序运行路径
func getCurrRunDir() string{
	dir,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	} 
	return strings.Replace(dir,"\\","/",-1)
}

func GetStringValue(key string)string{
	return viper.GetString(key)
}

func GetDbConectInfo()(*DbConnectInfo,error){
	err := viper.ReadInConfig()
	if err != nil{
		return nil,err
	}
	dbInfo := &DbConnectInfo{
					Ip:viper.GetString("mysql.host"),
					Port:viper.GetInt("mysql.port"),
					Username:viper.GetString("mysql.username"),
					Password:viper.GetString("mysql.password"),
					Dbname:viper.GetString("mysql.dbname"),
				}
	return dbInfo,nil
}

//
func GetCoinRpc(currency string)(*RpcConnectInfo, error){
	err := viper.ReadInConfig()
	if err != nil{
		return nil,err
	}
	var rpcInfo *RpcConnectInfo
	if currency=="eth" || currency=="etc"  {
		rpcInfo = &RpcConnectInfo{
			Currency:currency,
			Ip:viper.GetString(fmt.Sprintf("%s.ip",currency)),
			Port:viper.GetInt(fmt.Sprintf("%s.port",currency)),
			Username:"",
			Password:"",
		}
	}else{
		rpcInfo = &RpcConnectInfo{
						Currency:currency,
						Ip:viper.GetString(fmt.Sprintf("%s.ip",currency)),
						Port:viper.GetInt(fmt.Sprintf("%s.port",currency)),
						Username:viper.GetString(fmt.Sprintf("%s.rpcusername",currency)),
						Password:viper.GetString(fmt.Sprintf("%s.rpcpassword",currency)),
					}
	}
	return rpcInfo,nil
}

//
func GetCurrPath() string{
	return "D:/gits/coins/config/config.json"
}

//根据币种从json配置文件中获取相应rpc信息
func GetRpcInfo(currency string) (*RpcConnectInfo, error){
	// jsonfile := GetCurrRunDir()+"/config/config.json"
	jsonfile := GetCurrPath()
	filedata,err := io.ReadFile(jsonfile)
	if err != nil{
		return nil,err
	}
	var jsondata []RpcConnectInfo
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