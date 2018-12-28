package models

import (
	math_rand "math/rand"
    "encoding/base64"
    "encoding/hex"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
	"errors"
	"time"
)


func GetPhraseAndSecret(pwd string,count int)(string, string){
	phrases := []string{"encodings","character","different","database","multiple","atomically","iterator","readrandom","repository","contents","information","below","accompanied","changes","submitting","leveldb","contributing","building","limitations","documentation","proof","supported","lighter","custom","fully","equivalent","network","reconfigure","instance","developers","around","creating","contracts","almost","certainly","involved","until","entire","towards","full","catch","hold","hope","city","software","big","buf","prime","parse","black","manager","computer","runner","terminal","edit","selection","view"}
	var key string
	p_len := len(phrases)
	rad := math_rand.New(math_rand.NewSource(time.Now().Unix()))
	for i:=0;i<count;i++{
		j := rad.Intn(p_len)
		key = phrases[j]+" " + key
    }
    key += "a very private public secret key for me"
    key = key[:32]
    result,err := Encrypt([]byte(key),[]byte(pwd))
    if err == nil{
        return key,hex.EncodeToString(result)
    }
	return "",""
}


func GetRealPwd(key string ,pwd string)(string,error){
    newPwd,err := hex.DecodeString(pwd)
    if err == nil{
        result,err1:=Decrypt([]byte(key),newPwd)
        return string(result),err1
    }
    return "",nil
}


func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}


func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}