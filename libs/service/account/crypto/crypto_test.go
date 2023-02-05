package crypto_service_test

import (
	"fmt"
	"testing"

	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func TestCrypto(t *testing.T) {
	content := "cipher content"
	cipherText, err := crypto_service.EncodeString(content)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	fmt.Println("cipherText -> ", cipherText)

	decodeText, err := crypto_service.DecodeString(cipherText)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	if decodeText != content {
		t.Errorf("解密后的字符串与初始化的字符串不相同：%s\t%s", content, decodeText)
		return
	}
}
