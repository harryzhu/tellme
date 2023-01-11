package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
)

// PwdKey length can be 16,24,32

var PwdKey = []byte(MD5(GetEnv("HAZHUENCRYPTKEY", "This(*Key*)@2021This(*Key*)@2021")))

// ------------

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("encryt text error")
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)], nil

}

func aesEnCrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs7Padding(origData, blockSize)
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	blockMode.CryptBlocks(origData, cypted)
	origData, err = pkcs7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

func AESEncode(pwd []byte) (string, error) {
	result, err := aesEnCrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

func AESDecode(pwd string) ([]byte, error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	return aesDeCrypt(pwdByte, PwdKey)

}

func GetEnv(s string, vDefault string) string {
	v := os.Getenv(s)
	if v == "" {
		return vDefault
	}
	return v
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	dst, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(dst)
}
