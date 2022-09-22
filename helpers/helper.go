package helper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/egeaybars123/randomethtx"
	ethcrypto "github.com/egeaybars123/randomethtx/cryptotx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func SendTx(privKey *ecdsa.PrivateKey, client *ethclient.Client, value *big.Int) error {
	_, toaddress := ethcrypto.GenerateAddress()
	sender := common.HexToAddress(randomethtx.SenderAddress)
	//fmt.Println(sender)
	//fmt.Println(privKey)

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
	tx := types.NewTransaction(nonce, toaddress, value, 500000, gasprice, nil)
	signed, _ := types.SignTx(tx, types.NewEIP155Signer(chainid), privKey)

	return client.SendTransaction(context.Background(), signed)
}

func RPCClient() (*rpc.Client, error) {
	cl, err := rpc.Dial(randomethtx.RpcAddress)

	return cl, err
}
