package transactions

import (
	"fmt"
	"math/big"

	ethcrypto "github.com/egeaybars123/randomethtx/cryptotx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTx(client *ethclient.Client, to *common.Address, value *big.Int) {
	privKey, sender := ethcrypto.GenerateAddress()
	fmt.Println(sender)
	fmt.Println(privKey)
}
