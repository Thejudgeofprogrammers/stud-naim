package impl

import (
	"errors"
	"time"

	JWTService "gateway/internal/service/jwt"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secret string
	exp    int
}

func NewJWTService(secret string, exp int) JWTService.JWTService {
	return &jwtService{secret: secret, exp: exp}
}

func (s *jwtService) GenerateAccessToken(userID string, role JWTService.AccessRole) (string, error) {
	claims := JWTService.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.exp) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *jwtService) Parse(tokenStr string) (*JWTService.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTService.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTService.Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
