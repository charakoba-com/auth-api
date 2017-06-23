package utils

import (
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/charakoba-com/auth-api/keymgr"
	"github.com/pkg/errors"
)

// GenerateToken generates a JSON Web Token
func GenerateToken(username string, isAdmin bool) (string, error) {
	claims := jws.Claims{}
	now := time.Now()
	expiration := now.Add(time.Duration(168) * time.Hour)
	claims.Set("username", username)
	claims.Set("is_admin", isAdmin)
	claims.SetExpiration(expiration)

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)
	privateKey, err := keymgr.PrivateKey()
	if err != nil {
		return "", errors.Wrap(err, `loading private key`)
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		return "", errors.Wrap(err, `serialize token`)
	}
	return string(token), nil
}
