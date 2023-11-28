package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// RSAPadding RSA PEM 填充模式
type RSAPadding int

const (
	RSA_PKCS1 RSAPadding = 1 // PKCS#1 (格式：`RSA PRIVATE KEY` & `RSA PUBLIC KEY`)
	RSA_PKCS8 RSAPadding = 8 // PKCS#8 (格式：`PRIVATE KEY` & `PUBLIC KEY`)
)

// GenerateRSAKey 生成RSA私钥和公钥
func GenerateRSAKey(bitSize int, padding RSAPadding) (privateKey, publicKey []byte, err error) {
	prvKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return
	}

	switch padding {
	case RSA_PKCS1:
		privateKey = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(prvKey),
		})

		publicKey = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&prvKey.PublicKey),
		})
	case RSA_PKCS8:
		prvBlock := &pem.Block{
			Type: "PRIVATE KEY",
		}

		prvBlock.Bytes, err = x509.MarshalPKCS8PrivateKey(prvKey)
		if err != nil {
			return
		}

		pubBlock := &pem.Block{
			Type: "PUBLIC KEY",
		}

		pubBlock.Bytes, err = x509.MarshalPKIXPublicKey(&prvKey.PublicKey)
		if err != nil {
			return
		}

		privateKey = pem.EncodeToMemory(prvBlock)
		publicKey = pem.EncodeToMemory(pubBlock)
	}

	return
}
