package crypto_service

import "encoding/hex"

// 加密
func EncodeString(content string) string {
	k := MustGetRsaPair()

	cipherBuf := RsaEncrypt([]byte(content), []byte(k.PubKey))
	return hex.EncodeToString(cipherBuf)
}

// 解密
func DecodeString(cipherText string) string {
	k := MustGetRsaPair()

	buf, _ := hex.DecodeString(cipherText)
	content := RsaDecrypt(buf, []byte(k.PrvKey))
	return string(content)
}
