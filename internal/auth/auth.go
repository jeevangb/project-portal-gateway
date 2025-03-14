package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var privateKey *rsa.PrivateKey

var publicKey *rsa.PublicKey

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   email,
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func LoadPrivateKey(path string) error {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return errors.New("invalid private key")
	}
	privateKeyinterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	parsedKey, ok := privateKeyinterface.(*rsa.PrivateKey)
	if !ok {
		return errors.New("private key is not an RSA key")
	}
	privateKey = parsedKey
	return nil
}

// LoadPublicKey loads the public key from a PEM file
func LoadPublicKey(publicKeyPath string) error {
	keyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return errors.New("invalid public key")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	var ok bool
	publicKey, ok = pubKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("not an RSA public key")
	}
	return nil
}
