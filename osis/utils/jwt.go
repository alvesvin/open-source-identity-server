package utils

import (
	"crypto/ed25519"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jose-util/generator"
)

func GeneratePrivateKey() {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	jwk := jose.JSONWebKey{}
}

func GetPrivateKey() {
}
