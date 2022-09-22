package transactions

import (
	"context"
	"fmt"
	"math/big"

	ethcrypto "github.com/egeaybars123/randomethtx/cryptotx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTx(client *ethclient.Client, to common.Address, value *big.Int) error {
	privKey, sender := ethcrypto.GenerateAddress()
	fmt.Println(sender)
	fmt.Println(privKey)

	//retrieves the pending nonce for a sender
	//this nonce should be used for the next transaction from the sender
	nonce, err := client.PendingNonceAt(context.Background(), sender)

	if err != nil {
		fmt.Printf("No nonce retrieved: %v \n", err)
	}
	fmt.Printf("Nonce for %v: %v \n", sender, nonce)
	chainid, err := client.ChainID(context.Background())

	if err != nil {
		return err
	}

	gasprice, _ := client.SuggestGasPrice(context.Background())
	tx := types.NewTransaction(nonce, to, value, 500000, gasprice, nil)
	signed, _ := types.SignTx(tx, types.NewEIP155Signer(chainid), privKey)

	return client.SendTransaction(context.Background(), signed)

}
