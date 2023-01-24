package login_service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/exceptions/exception"
	file_service "github.com/sheason2019/spoved/libs/service/file"
	"github.com/sheason2019/spoved/libs/utils"
)

var jwt_secret_file = "jwt_secret"

var jwtSecret = GenerateJwtSecret()

type JwtClaims struct {
	ent.User
	jwt.RegisteredClaims
}

// 生成
func GenerateJwt(user *ent.User) (string, *exception.Exception) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "Spoved",
		},
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", exception.New(err)
	}
	return tokenString, nil
}

// 解析
func ParseJwt(tokenString string) (*JwtClaims, *exception.Exception) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, exception.New(err)
}

// 生成Jwt的加密秘钥
func GenerateJwtSecret() string {
	// 首先尝试从文件系统中获取
	jwtSecret, err := file_service.Read(jwt_secret_file)
	if err == nil && len(jwtSecret) > 0 {
		return jwtSecret
	}

	// 否则以64位随机字符串作为jwt的加密秘钥
	jwtSecret = utils.RandomStr(64)
	// 并存入文件系统
	err = file_service.Write(jwtSecret, jwt_secret_file)
	if err != nil {
		exception.New(err).Panic()
	}

	return jwtSecret
}
