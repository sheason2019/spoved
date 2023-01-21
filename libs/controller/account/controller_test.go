package account_controller_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/router"
	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func TestRegist(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", account.AccountServiceDefinition.ACCOUNT_CRYPTO_PATH, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)

	var cryptoInfo account.AccountCrypto
	err := json.Unmarshal(w.Body.Bytes(), &cryptoInfo)
	if err != nil {
		t.Error(err)
		return
	}

	testUsername := "testUser1"
	testPassword := "testPassword"

	cipherPasswordBuf := crypto_service.RsaEncrypt(
		[]byte(testPassword+cryptoInfo.Salt),
		[]byte(cryptoInfo.PubKey),
	)

	cipherPassword := hex.EncodeToString(cipherPasswordBuf)

	info := account.AccountInfo{
		Username: testUsername,
		Password: cipherPassword,
		Salt:     cryptoInfo.Salt,
	}
	buf, _ := json.Marshal(info)
	body := bytes.NewReader(buf)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", account.AccountServiceDefinition.REGIST_PATH, body)
	router.ServeHTTP(w, req)

	// testUser1已存在
	assert.NotEqual(t, w.Code, 200)
}
