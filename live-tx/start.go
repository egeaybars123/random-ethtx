package main

import (
	"fmt"

	crypto "github.com/egeaybars123/randomethtx/cryptotx"
)

func main() {
	priv, err := crypto.GeneratePrivKey()
	keystr := crypto.KeyToString(priv)
	//pub := tx.ShowPubKey(priv)
	//fmt.Println(reflect.TypeOf(priv))
	address := crypto.KeyToAddress(&priv.PublicKey)
	fmt.Println(priv, err)
	fmt.Println(address)
	fmt.Println(keystr)
}
