package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	RSA_DIR_NAME    = "dcc"
	RSA_PUBLIC_KEY  = "public_key.pem"
	RSA_PRIVATE_KEY = "private_key.pem"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func SavePrivateEMKey(filename string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.Encode(outFile, privateKeyPEM)
}

func SavePublicPEMKey(filename string, pubkey *rsa.PublicKey) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return pem.Encode(outFile, publicKeyPEM)
}

func ReadPrivateKey(filename string) ([]byte, error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key PEM data")
	}
	return block.Bytes, nil
}

func ReadPublicKey(filename string) ([]byte, error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key PEM data")
	}

	return block.Bytes, nil
}

func EncryptData(filename string, plaintext []byte) ([]byte, error) {
	publicKeyData, err := ReadPublicKey(filename)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(publicKeyData)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub.(*rsa.PublicKey), plaintext, nil)
}

func DecryptData(filename string, ciphertext []byte) ([]byte, error) {
	privateKeyData, err := ReadPrivateKey(filename)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyData)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
}

func DecryptAES256CFB(encryptedData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(encryptedData))
	stream.XORKeyStream(plaintext, encryptedData)

	return plaintext, nil
}
