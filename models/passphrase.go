package models

import (
	math_rand "math/rand"
    "encoding/base64"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "errors"
)


func GetPhraseAndSecret(pwd string,count int)(string, []byte){
	phrases := []string{"repository","contents","information","below","accompanied","changes","submitting","leveldb","contributing","building","limitations","documentation","proof","supported","lighter","custom","fully","equivalent","network","reconfigure","instance","developers","around","creating","contracts","almost","certainly","involved","until","entire","towards","full","catch","hold","hope","city","software","big","buf","prime","parse","black","manager","computer","runner","terminal","edit","selection","view"}
	var key string
	p_len := len(phrases)
	for i:=0;i<count;i++{
		j := math_rand.Intn(p_len)
		key = phrases[j]+" " + key
    }
    key += "a very private public secret key for me"
    key = key[:32]
    result,err := Encrypt([]byte(key),[]byte(pwd))
    if err == nil{
        return key,result
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
