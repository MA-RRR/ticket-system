package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UID  uint
	Role string
	jwt.RegisteredClaims
}

func GenerateToken(uid uint, role string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UID:  uid,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwt.TimeFunc().Add(time.Hour * 2)),
		},
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
