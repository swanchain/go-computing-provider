package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type Signature struct {
	Data []byte
}

// Sign takes in signature type, private key and message. Returns a signature for that message.
func Sign(privatekey string, msg []byte) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256Hash(msg)
	return crypto.Sign(hash.Bytes(), privateKey)
}

// Verify verifies signatures
func Verify(addr string, signature, dataHash []byte) (bool, error) {
	sigPublicKeyECDSA, err := crypto.SigToPub(dataHash, signature)
	if err != nil {
		return false, err
	}
	pub := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
	if addr != pub {
		return false, err
	}
	return true, nil
}

// ToPublic converts private key to public key
func ToPublic(priv string) (string, *ecdsa.PublicKey, error) {
	if priv == "" || len(strings.TrimSpace(priv)) == 0 {
		return "nil", nil, fmt.Errorf("invalid private key")
	}

	privateKeyBytes, err := hex.DecodeString(priv)
	if err != nil {
		return "", nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicK := hexutil.Encode(publicKeyBytes)[4:]
	return publicK, publicKeyECDSA, nil
}
