package main

import (
	"math/big"
	"time"

	log "github.com/sirupsen/logrus"

	tx "github.com/egeaybars123/randomethtx"
	helper "github.com/egeaybars123/randomethtx/helpers"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func publishTx() {
	//priv, err := crypto.GeneratePrivKey()
	//keystr := crypto.KeyToString(priv)
	//address := crypto.KeyToAddress(&priv.PublicKey)
	backend, _ := helper.RPCClient()
	client := ethclient.NewClient(backend)
	value := big.NewInt(1000000000000000000) //1 ether
	privKey := ethcrypto.ToECDSAUnsafe(common.FromHex(tx.SenderKey))

	_, to := helper.SendTx(privKey, client, value)

	log.WithFields(log.Fields{
		"to":    to,
		"value": value,
	})

	//fmt.Println(priv, err)
	//fmt.Println(address)
	//fmt.Println(keystr)
}

func main() {
	//pub := tx.ShowPubKey(priv)
	//fmt.Println(reflect.TypeOf(priv))

	log.Info("Starting the transaction spam...")

	ticker := time.NewTicker(500 * time.Millisecond)

	for range ticker.C {
		publishTx()
	}

}
