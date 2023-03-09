package jwt

import "time"
import jwtgo "github.com/dgrijalva/jwt-go"

var jwtSecret []byte

type Config struct {
	Secret string
	TTL    time.Duration
}

type Claims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   int32  `json:"sub,omitempty"`
	jwtgo.MapClaims
}

func Init(c *Config) {
	jwtSecret = []byte(c.Secret)
}

func GenerateToken(claims Claims) (string, error) {
	tokenClaims := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwtgo.ParseWithClaims(token, &Claims{}, func(token *jwtgo.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
