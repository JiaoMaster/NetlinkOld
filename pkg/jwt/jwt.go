package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var TokenExpireDuration = time.Hour * 7200

var mySecret = []byte("jiaomaster")

// GenToken 生成JWT
func GenToken(userid int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userid,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "jiao",                                     // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
