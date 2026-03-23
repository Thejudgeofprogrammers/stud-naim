package jwt

type JWTService interface {
	GenerateAccessToken(userID string, role AccessRole) (string, error)
	Parse(tokenStr string) (*Claims, error)
}
