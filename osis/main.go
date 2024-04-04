package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"

	"github.com/go-jose/go-jose/v4"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Fprint(w, err)
	}

	publicKey := &privateKey.PublicKey
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: publicKey}, nil)

	if err != nil {
		fmt.Fprint(w, err)
	}

	plaintext := []byte("Lorem ipsum dolor sit amet")
	object, err := encrypter.Encrypt(plaintext)

	if err != nil {
		fmt.Fprint(w, err)
	}

	serialized := object.FullSerialize()

	object, err = jose.ParseEncrypted(serialized, []jose.KeyAlgorithm{jose.RSA_OAEP}, []jose.ContentEncryption{jose.A128GCM})

	if err != nil {
		fmt.Fprint(w, err)
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)

	if err != nil {
		fmt.Fprint(w, err)
	}

	object2, err := signer.Sign([]byte("test"))
	serialized2 := object2.FullSerialize()
	if serialized2 != "" {
	}

	if err != nil {
		fmt.Fprint(w, err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, serialized)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
