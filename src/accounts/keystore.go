package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
	"os"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	file := "./tmp/UTC--2023-02-28T22-56-09.926500100Z--197a9049e40cbb7d6a3d7ba8a2f4e51afcffbf0b"
	ks := keystore.NewKeyStore("./tmp2", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	//if err := os.Remove(file); err != nil {
	//	log.Fatal(err)
	//}
}

func main() {
	//createKs()
	importKs()
}
