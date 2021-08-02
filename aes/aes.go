package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"

	"github.com/zengzhengrong/zzgo/curl"
)

type AesCrypt struct {
	Key []byte
	Iv  []byte
}

// aes加密
func (a *AesCrypt) Encrypt(data []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher(a.Key)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	content := pKCS5Padding(data, aesBlockEncrypt.BlockSize())
	cipherBytes := make([]byte, len(content))
	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, a.Iv)
	aesEncrypt.CryptBlocks(cipherBytes, content)
	return cipherBytes, nil
}

// aes解密
func (a *AesCrypt) Decrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err = aes.NewCipher(a.Key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, a.Iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return pKCS5Trimming(decrypted), nil
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// AesPost is http post 请求加密交互
// key,iv: 加密的key和偏移量iv
// aesdata,aeskey, 要加密的内容和对应键名
// url 要请求的地址
func AesPost(key, iv, aesdata, aeskey, url string) ([]byte, error) {

	crypto := AesCrypt{
		Key: []byte(key),
		Iv:  []byte(iv),
	}
	encrypt, err := crypto.Encrypt([]byte(aesdata))
	if err != nil {
		return nil, err
	}
	postData := map[string]string{
		aeskey: base64.StdEncoding.EncodeToString(encrypt),
	}
	postDataByte, _ := json.Marshal(postData)
	resp, err := curl.Post(url, postDataByte, nil)
	if err != nil {
		return nil, err
	}

	pass, err := base64.StdEncoding.DecodeString(string(resp))
	if err != nil {
		return nil, err
	}

	decrypt, err := crypto.Decrypt(pass)

	if err != nil {
		return nil, err
	}
	return decrypt, nil
}
