package auth

import (
	"time"
)

type Auth struct {
	jwtTTTL   time.Duration
	userMap   map[string]uint64
	secretKey []byte
	AuthActions
}

func NewAuth(jwtTTL int) Auth {
	return Auth{jwtTTTL: time.Duration(jwtTTL) * time.Minute, userMap: make(map[string]uint64), secretKey: []byte("SbEvE__52--a")}
}
