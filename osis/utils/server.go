package utils

import (
	"os"
)

func EnsurePrivateKey() {
	privateKeyPath := os.Getenv("OSIS_PRIVATE_KEY_PATH")
	if privateKeyPath == "" {
		privateKeyPath = "./"
	}
	GeneratePrivateKey()
}
