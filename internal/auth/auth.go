package auth

import (
	"time"

	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

var (
	JWTSecretKey = []byte("SbEvE__52--a")
)

type Auth struct {
	jwtTTTL   time.Duration
	userMap   map[string]model.AuthAns
	secretKey []byte
	AuthActions
}

func NewAuth(jwtTTL int) Auth {
	return Auth{jwtTTTL: time.Duration(jwtTTL) * time.Minute, userMap: make(map[string]model.AuthAns), secretKey: JWTSecretKey}
}
