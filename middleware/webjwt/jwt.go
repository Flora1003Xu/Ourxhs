//jwt应该好了，但是还没有测试

package webjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// 分发token
func ReleaseToken(phone string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 设置过期时间（左边为7天）
	claims := &Claims{
		UserId: phone, // 用户ID？
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), // token发放的时间
			Issuer:    "Ourxhs",          // token发放者
			Subject:   "user token",      // token主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HS256加密方式
	tokenString, err := token.SignedString(jwtKey)             // 使用秘钥jwtKey来生成token

	if err != nil {
		return "", err
	}

	return tokenString, nil // 返回生成的token
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
