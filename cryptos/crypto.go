package cryptos

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

// MD5 字符串
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// EncodeAuthHeader 加密
func EncodeAuthHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

// DecodeAuthHeader 解密
func DecodeAuthHeader(s string) (user, pass string, err error) {
	authData := strings.Split(s, " ")[1]

	var data []byte
	data, err = base64.StdEncoding.DecodeString(authData)
	if err != nil {
		return
	}
	list := strings.Split(string(data), ":")
	if len(list) != 2 {
		err = errors.New("invalid auth header")
		return
	}
	user, pass = list[0], list[1]
	return
}

// PKCS7Padding padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding unpadding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt AES加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// AesEncryptBase64 AES加密后Base64编码
func AesEncryptBase64(origData, key string) (string, error) {
	result, err := AesEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// AesDecryptBase64 Base64解码后Aes解密
func AesDecryptBase64(s, key string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	originData, err := AesDecrypt(bs, []byte(key))
	if err != nil {
		return "", err
	}

	return string(originData), nil
}
