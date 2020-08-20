package service

import (
	"HackInBox.online/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// SecretKey 认证key
const (
	SecretKey = "fdsafjkldsaklfjkdlasjfkljsdaklfjskdlafklsdjaklfjaslfjiouwiotfdsafasdfjikljsngklnsadvhasjkfghejskd"
	Issuer    = "bugbug"
)

// jwtCustomClaims token签名信息
type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid interface{}
}

// GenerateToken 生成token
func GenerateToken(Uid interface{}) string {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    Issuer,
		},
		Uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		utils.UtilsLogger.Error(err)
		return ""
	}
	return tokenString
}

// VerifyToken 验证token
func VerifyToken(tokenSrt string) jwt.Claims {
	//var verifyToken *jwt.Token
	verifyToken, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil
	}
	claims := verifyToken.Claims
	return claims
}
