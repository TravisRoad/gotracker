package auth_test

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	expTime := jwt.NewNumericDate(time.Now().Add(-5 * time.Minute))

	claims := jwt.MapClaims{
		"exp": expTime,
		"nbf": jwt.NewNumericDate(time.Now()),
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("foo"))
	if err != nil {
		panic(err)
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("foo"), nil
	})

	assert.NotEqual(t, err, nil)
}
