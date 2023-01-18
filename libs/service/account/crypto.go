package account_service

// 加密
func EncodeString(content string) string {
	k := MustGetRsaPair()

	cipherBuf := RsaEncrypt([]byte(content), []byte(k.PubKey))
	return string(cipherBuf)
}

// 解密
func DecodeString(cipherText string) string {
	k := MustGetRsaPair()

	content := RsaDecrypt([]byte(cipherText), []byte(k.PrvKey))
	return string(content)
}
