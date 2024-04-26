package jwtV1

import "github.com/golang-jwt/jwt/v5"

type JwtValidator interface {
	Encode(claims jwt.Claims) (string, error)
	Decode(sign string, claims jwt.Claims) error
}
