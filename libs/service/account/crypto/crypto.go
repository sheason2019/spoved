package crypto_service

import (
	"encoding/base64"
	"encoding/hex"
	"errors"

	"github.com/sheason2019/spoved/exceptions/exception"
)

// 加密
func EncodeString(content string) string {
	k := MustGetRsaPair()

	cipherBuf := RsaEncrypt([]byte(content), []byte(k.PubKey))
	return hex.EncodeToString(cipherBuf)
}

// 解密
func DecodeString(cipherText string) (string, *exception.Exception) {
	k := MustGetRsaPair()

	cipherPassword, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", exception.New(errors.New("密文解析失败"))
	}
	buf := []byte(cipherPassword)
	content, e := RsaDecrypt(buf, []byte(k.PrvKey))
	if e != nil {
		return "", e.Wrap()
	}
	return string(content), nil
}
