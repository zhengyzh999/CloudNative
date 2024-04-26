package jwtV1

import (
	"github.com/golang-jwt/jwt/v5"
)

type RS struct {
	PublicKey  string
	PrivateKey string
}

func (rs *RS) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(rs.PrivateKey))
	if err != nil {
		return "", nil
	}
	sign, err := token.SignedString(privateKey)
	return sign, nil
}
func (rs *RS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(rs.PublicKey))
	})
	return err
}
