package account_service_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	account_service "github.com/sheason2019/spoved/libs/service/account"
)

func TestCrypto(t *testing.T) {
	content := "cipher content"
	cipherText := account_service.EncodeString(content)
	fmt.Println("cipherText -> ", hex.EncodeToString([]byte(cipherText)))

	decodeText := account_service.DecodeString(cipherText)
	if decodeText != content {
		t.Errorf("解密后的字符串与初始化的字符串不相同：%s\t%s", content, decodeText)
	}
}
