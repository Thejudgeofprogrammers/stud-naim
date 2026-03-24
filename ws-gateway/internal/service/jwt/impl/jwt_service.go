package impl

import (
	JWTService "ws-gateway/internal/service/jwt"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secret string
	exp    int
}

func NewJWTService(secret string, exp int) JWTService.JWTService {
	return &jwtService{secret: secret, exp: exp}
}

func (s *jwtService) Parse(tokenStr string) (*JWTService.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTService.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*JWTService.Claims)
	return claims, nil
}