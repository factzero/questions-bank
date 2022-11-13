package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type BaseClaimsRequest struct {
	UUID        uuid.UUID
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
