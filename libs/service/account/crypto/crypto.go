package crypto_service

import (
	"encoding/hex"

	"github.com/pkg/errors"
)

// 加密
func EncodeString(content string) (string, error) {
	k := MustGetRsaPair()

	cipherBuf, err := RsaEncrypt([]byte(content), []byte(k.PubKey))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherBuf), nil
}

// 解密
func DecodeString(cipherText string) (string, error) {
	k := MustGetRsaPair()

	cipherPassword, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", errors.WithStack(errors.New("密文解析失败"))
	}

	content, e := RsaDecrypt(cipherPassword, []byte(k.PrvKey))
	if e != nil {
		return "", e
	}
	return string(content), nil
}
