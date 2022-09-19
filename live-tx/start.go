package main

import (
	"fmt"

	tx "github.com/egeaybars123/randomethtx/cryptotx"
)

func main() {
	a, b := tx.GeneratePrivKey()
	fmt.Println(a, b)
}
