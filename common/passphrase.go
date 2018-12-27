package common

import (
	"math/rand"
	"crypto/aes"
	"bytes"
	"crypto/cipher"
)


func GetPhraseAndSecret(pwd string,count int)(string, string,error){
	phrases := []string{"proof","supported","lighter","custom","fully","equivalent","network","reconfigure","instance","developers","around","creating","contracts","almost","certainly","involved","until","entire","towards","full","catch","hold","hope","city","software","big","buf","prime","parse","black","manager","computer","runner","terminal","edit","selection","view"}
	var result string 
	for i:=0;i<count;i++{
		j := rand.Intn(count)
		result = phrases[j]+" " + result
	}
	secret,err := AesEncrypt([]byte(pwd),[]byte(result))
	if err != nil{
		return "","",err
	}
	return result,string(secret),nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}

