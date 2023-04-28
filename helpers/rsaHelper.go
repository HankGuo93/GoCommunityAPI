package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	var privateKey *rsa.PrivateKey
	privateKeyPEMBytes, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		return privateKey, err
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEMBytes)
	if privateKeyBlock == nil {
		return privateKey, errors.New("failed to decode PEM block containing private key")
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return privateKey, errors.New(err.Error())
	}
	return privateKey, err
}

func GetPublicKey() (interface{}, error) {
	publicKeyPEMBytes, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		return nil, err
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEMBytes)
	if publicKeyBlock == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return publicKey, err
}
