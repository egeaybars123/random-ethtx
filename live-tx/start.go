package main

import (
	"fmt"

	tx "github.com/egeaybars123/randomethtx/cryptotx"
)

func main() {
	priv, err := tx.GeneratePrivKey()
	//pub := tx.ShowPubKey(priv)
	//fmt.Println(reflect.TypeOf(priv))
	pub := tx.ShowPubKey(priv)
	address := tx.KeyToAddress(&priv.PublicKey)
	fmt.Println(priv, err)
	fmt.Println(pub)
	fmt.Println(address)
}
