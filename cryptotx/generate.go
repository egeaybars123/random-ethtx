package cryptotx

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func GeneratePrivKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	return key, err
}

func ShowPubKey(privKey *ecdsa.PrivateKey) crypto.PublicKey {
	return privKey.Public()
}

func KeyToAddress(pubKey *ecdsa.PublicKey) common.Address {
	pubBytes := elliptic.Marshal(secp256k1.S256(), pubKey.X, pubKey.Y)
	return common.BytesToAddress(ethcrypto.Keccak256(pubBytes[1:])[12:])
}
