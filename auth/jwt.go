package auth

import (
	"github.com/dgrijalva/jwt-go"
)

const JwtSecretKey = "fastChat-jwt"

type JwtClaims struct {
	UserId string

	OpenId string

	Namespace string

	Exp int64 //过期时间
}

//CreateJwtToken 创建JWT字符串
func CreateJwtToken(claims JwtClaims) (string, error) {

	rawClaims := jwt.MapClaims{
		"exp": claims.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rawClaims)

	t, e := token.SignedString([]byte(JwtSecretKey))

	return t, e
}

//ParseJwtToken 校验jwt字符串
func ParseJwtToken(s string) (*JwtClaims, error) {

	fn := func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil //校验字符串
	}

	result, error := jwt.Parse(s, fn)
	if error != nil {
		return nil, error //signature is invalid or Token is expired
	}

	//解析存入去的jwt信息
	rawClaims := result.Claims.(jwt.MapClaims)

	return &JwtClaims{
		Exp: rawClaims["exp"].(int64),
	}, nil
}
