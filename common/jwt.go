package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tzk-chris/Gin-Vue/model"
	"time"
)

var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // token保存时间一周
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}

// postman测试：
//{
//"code": 200,
//"data": {
//"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImV4cCI6MTYxODgwNDE5NiwiaWF0IjoxNjE4MTk5Mzk2LCJzdWIiOiJ1c2VyIHRva2VuIn0.Gc3wZ0Y8Z0N4fPqAlKD2GZrZ2ddN3shy7C79f2jS8yw"
//},
//"msg": "登录成功"
//}

// token组成：
//1. 协议头  -- 这里的协议是 jwt
//2. 定义的claims数据
//3. 前面两个部分加上定义的密钥的hash值
