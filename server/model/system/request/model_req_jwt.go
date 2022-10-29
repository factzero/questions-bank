package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type BaseClaimsRequest struct {
	UUID        string
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}

type CustomClaimsRequest struct {
	BaseClaimsRequest
	BufferTime int64
	jwt.StandardClaims
}
