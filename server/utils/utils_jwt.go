package utils

import (
	"errors"
	"server/global"
	systemReq "server/model/system/request"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (jn *JWT) CreateClaims(baseClaims systemReq.BaseClaimsRequest) systemReq.CustomClaimsRequest {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := systemReq.CustomClaimsRequest{
		BaseClaimsRequest: baseClaims,
		BufferTime:        int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,     // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),    // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}

	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims systemReq.CustomClaimsRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*systemReq.CustomClaimsRequest, error) {
	token, err := jwt.ParseWithClaims(tokenString, &systemReq.CustomClaimsRequest{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*systemReq.CustomClaimsRequest); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}
