package jwt

type JWTService interface {
	Parse(tokenStr string) (*Claims, error)
}
