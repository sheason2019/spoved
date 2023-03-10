package crypto_service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func MustGetRsaPair() (keyPair *RsaKeyPair) {
	k, e := GetRsaPair()
	if e == nil {
		return k
	}
	k, e = InitRsa()
	if e == nil {
		return k
	}
	panic(e)
}

// 取得密钥对
func GetRsaPair() (keyPair *RsaKeyPair, e error) {
	keyPair = &RsaKeyPair{}
	prvContent, err := file_service.Read(_key_path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	keyPair.PrvKey = string(prvContent)

	pubContent, err := file_service.Read(_pub_key_path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	keyPair.PubKey = string(pubContent)

	return
}

// 生成RSA密钥对文件
func InitRsa() (*RsaKeyPair, error) {
	// 若产生错误，则表示无法获取到当前的RSA密钥对，需要重新生成
	keyPair, err := genRsaKey()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 将KeyPair写入文件系统实现持久化
	err = keyPair.persist()
	if err != nil {
		return nil, err
	}
	return keyPair, nil
}

// 生成密钥对
func genRsaKey() (*RsaKeyPair, error) {
	keyPair := &RsaKeyPair{}
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey := pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey := pem.EncodeToMemory(block)
	keyPair.PrvKey = string(prvkey)
	keyPair.PubKey = string(pubkey)
	return keyPair, nil
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.WithStack(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ciphertext, nil
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.WithStack(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}

// 持久化存储密钥对
const _pub_key_path = "/rsa_key.pub"
const _key_path = "/rsa_key"

func (k *RsaKeyPair) persist() error {
	err := file_service.Write(k.PubKey, _pub_key_path)
	if err != nil {
		return errors.WithStack(err)
	}
	err = file_service.Write(k.PrvKey, _key_path)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
