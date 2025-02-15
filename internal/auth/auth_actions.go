package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

type AuthActions interface {
	GenerateFromPassword(model.IdPassword) (string, error)
	Identify(token string) (uint64, error)
}

func (auth *Auth) GenerateFromPassword(data model.IdPassword) (string, error) {

	expirationTime := time.Now().Add(auth.jwtTTTL)

	claims := &Claims{
		UserID: data.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(auth.secretKey)
	if err != nil {
		return "", err
	}

	tokenHash := sha256Hash(tokenString)
	auth.userMap[tokenHash] = data.Id

	return tokenString, nil
}
func (auth *Auth) Identify(token string) (val uint64, ok bool) {
	val, ok = auth.userMap[token]
	return
}
