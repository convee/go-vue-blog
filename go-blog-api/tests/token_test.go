package tests

import (
	"github.com/convee/go-blog-api/pkg/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	c := &jwt.Config{
		Secret: "CD1B3SVGZOJ0dR4j7cML2mvoKkePqrUn",
	}
	jwt.Init(c)
	claim := jwtgo.StandardClaims{
		Subject:   "convee",
		ExpiresAt: time.Now().Unix() + 31104000,
	}
	token, err := jwt.GenerateToken(claim)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parseToken)
}
