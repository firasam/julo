package service

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(xID string) (string, error)
}

type jwtService struct {
}

type JWTSConfig struct {
}

func NewJWTService(c *JWTSConfig) JWTService {
	return &jwtService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET"))
var JWT_TTL, _ = strconv.Atoi(os.Getenv("JWT_TTL"))

type idTokenClaims struct {
	jwt.RegisteredClaims
	Xid string `json:"xid"`
}

func (s *jwtService) GenerateToken(xID string) (string, error) {
	payload := idTokenClaims{}
	payload.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Minute * time.Duration(JWT_TTL))}
	payload.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	payload.Xid = xID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
