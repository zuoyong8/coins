//copyright 

package rpc

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
	"errors"
	"encoding/json"
	"bytes"
  )

type JsonRpc struct {
	Host string         	//rpc主机
	Port int           		//rpc端口
	User string   		 	//rpc用户名
	Password string     	//rpc密码
}


func New(host string, port int,user string, password string) *JsonRpc {
	c := JsonRpc{host, port,user, password}
	return &c
}


//rpc认证连接，返回获取数据
func (jrpc *JsonRpc) MakeRequest(method string, params []interface{})([]byte, error)  {
		baseUrl := fmt.Sprintf("http://%s:%d", jrpc.Host, jrpc.Port)
		client := new(http.Client)
		req, err := http.NewRequest("POST", baseUrl, nil)
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(jrpc.User, jrpc.Password)
		req.Header.Add("Content-Type", "text/plain")
		
		args := make(map[string]interface{})
		args["jsonrpc"] = "2.0"
		args["id"] = time.Now().UnixNano()
		args["method"] = method
		args["params"] = params
		
		j, err := json.Marshal(args)
		if err != nil {
			return nil,err
		}
		
		req.Body = ioutil.NopCloser(strings.NewReader(string(j)))
		req.ContentLength = int64(len(string(j)))
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
}
		

//解析rpc方法返回的json数组
func RpcJsonParse(inDatas []byte)([]byte,error){
	var data  map[string]interface{}
	decoder  := json.NewDecoder(bytes.NewBuffer(inDatas))
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		str,_ := json.Marshal(err)
		return nil, errors.New(string(str))
	}

	err1 := data["error"];
	if err1 != nil {
		str,_ := json.Marshal(err1)
		return nil, errors.New(string(str))
	}

	result := data["result"]
	if result != nil {
		rbytes,err := json.Marshal(result)
		if err == nil {
			return rbytes, nil
		}
	}
	return nil, errors.New("no result")
}



