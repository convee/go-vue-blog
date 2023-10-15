package jwt

import "time"
import jwtgo "github.com/dgrijalva/jwt-go"

var jwtSecret []byte

type Config struct {
	Secret string
	TTL    time.Duration
}

func Init(c *Config) {
	jwtSecret = []byte(c.Secret)
}

func GenerateToken(claims jwtgo.StandardClaims) (string, error) {
	tokenClaims := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*jwtgo.StandardClaims, error) {
	tokenClaims, err := jwtgo.ParseWithClaims(token, &jwtgo.StandardClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtgo.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
