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
		User string   		  //rpc用户名
		Password string     //rpc密码
		Host string         //rpc主机
		Port int           //rpc端口
}
  

//rpc认证连接，返回获取数据
func (jrpc *JsonRpc) MakeRequest(method string, params []string)([]byte, error)  {
		baseUrl := fmt.Sprintf("http://%s:%d", jrpc.Host, jrpc.Port)
		client := new(http.Client)
		req, err := http.NewRequest("POST", baseUrl, nil)
		if err != nil {
			return nil, err
		}
		
		req.SetBasicAuth(jrpc.User, jrpc.Password)
		req.Header.Add("Content-Type", "text/plain")
		
		args := make(map[string]interface{})
		args["jsonrpc"] = "1.0"
		args["id"] = time.Now().UnixNano()
		args["method"] = method
		args["params"] = params
		
		j, err := json.Marshal(args)
		if err != nil {
			fmt.Println(err)
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
		

//解析非数组格式的json
func (c *JsonRpc) JsonParseToMapString(inDatas []byte)(map[string]interface{},error){
	var data map[string]interface{}
	decoder  := json.NewDecoder(bytes.NewBuffer(inDatas))
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		str,_ := json.Marshal(err)
		return nil, errors.New(string(str))
	}

	if err, found := data["error"]; found && err != nil {
		str,_ := json.Marshal(err)
		return nil, errors.New(string(str))
	}
		
	if result, found := data["result"]; found {
		return result.(map[string]interface{}), nil
	} else {
		return nil, errors.New("no result")
	}
}


//解析数组格式的json
func (c *JsonRpc) JsonParseToArray(inDatas []byte)([]interface{},error){
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

	result := data["result"].([]interface{})
	if result != nil {
		return result, nil
	}
	return nil, errors.New("no result")
}


func NewClient(user string, password string, host string, port int) *JsonRpc {
	c := JsonRpc{user, password, host, port}
	return &c
}