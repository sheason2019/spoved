package crypto_service_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func TestCrypto(t *testing.T) {
	content := "cipher content"
	cipherText := crypto_service.EncodeString(content)
	fmt.Println("cipherText -> ", hex.EncodeToString([]byte(cipherText)))

	decodeText, e := crypto_service.DecodeString(cipherText)
	if e != nil {
		e.Panic()
	}
	if decodeText != content {
		t.Errorf("解密后的字符串与初始化的字符串不相同：%s\t%s", content, decodeText)
	}
}
