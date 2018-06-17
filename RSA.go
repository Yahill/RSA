package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	option := flag.String("option", "waiting", "Choose option 1. keys - for generating keys 2.crypt - for crypting message 3. decrypt - for decrypting message")
	openKey := flag.Int64("open", 1, "Enter your open key")
	closedKey := flag.Int64("closed", 1, "Enter your closed key")
	signature := flag.Int64("signature", 1, "Enter your signature")

	flag.Parse()

	switch *option {

	case "keys":
		getKeys()

	case "crypt":
		crypt(*openKey, *signature)

	case "decrypt":
		decrypt(*closedKey, *signature)

	case "waiting":
		fmt.Println("Please choose what you want to do.")

	}
}
