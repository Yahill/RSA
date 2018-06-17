package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
)

type Data struct {
	message          []byte
	messageBig       []*big.Int
	cryptedMessage   []*big.Int
	decryptedMessage []*big.Int
	output           []byte
	buff             string
}

func crypt(openKey, signature int64) {
	data := Data{}

	data.message = readFile("message.txt")
	data.messageBig = byteIntoBigInt(data.message)
	data.cryptedMessage = cryptMessage(data.messageBig, openKey, signature)
	data.output = bigIntToByte(data.cryptedMessage)
	writeFile("crypted_message.txt", data.output)
	fmt.Println("Your message sucsefully crypted.")
}

func decrypt(closedKey, signature int64) {
	data := Data{}

	data.decryptedMessage = readCryptedMessage()
	data.decryptedMessage = decryptMessage(data.decryptedMessage, closedKey, signature)
	data.buff = bigIntToText(data.decryptedMessage)
	data.output = stringToByte(data.buff)
	writeFile("decrypted_message.txt", data.output)
	fmt.Println("Your message sucsefully decrypted")
}

func readFile(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return file
}

func writeFile(filePath string, data []byte) {
	ioutil.WriteFile(filePath, data, 0644)
}

func byteIntoBigInt(message []byte) []*big.Int {
	j := 1

	var messageIntoBigInt []*big.Int

	//convert []byte message into the []*big.Int message
	for i := 0; i < len(message); i++ {

		buff := new(big.Int)
		buff.SetBytes(message[i:j])
		j++

		messageIntoBigInt = append(messageIntoBigInt, buff)
	}
	return messageIntoBigInt
}

func bigIntToByte(file []*big.Int) (outputByte []byte) {
	output := fmt.Sprint(file)
	outputByte = []byte(output)
	return
}

func cryptMessage(messageBig []*big.Int, openKey, module int64) []*big.Int {

	for i := 0; i < len(messageBig); i++ {

		//power
		buff := (new(big.Int).Exp(messageBig[i], big.NewInt(openKey), nil))
		//%
		messageBig[i] = new(big.Int).Mod(buff, big.NewInt(module))

	}

	for i := 0; i < len(messageBig)-1; i++ {
		buff := (new(big.Int).Add(messageBig[i+1], messageBig[i]))
		messageBig[i+1] = new(big.Int).Mod(buff, big.NewInt(module))
	}

	return messageBig
}

func decryptMessage(messageBig []*big.Int, closedKey, module int64) []*big.Int {

	for i := len(messageBig) - 1; i > 0; i-- {
		buff := (new(big.Int).Sub(messageBig[i], messageBig[i-1]))
		messageBig[i] = new(big.Int).Mod(buff, big.NewInt(module))
	}

	for i := 0; i < len(messageBig); i++ {
		//power
		buff := (new(big.Int).Exp(messageBig[i], big.NewInt(closedKey), nil))
		//%
		messageBig[i] = new(big.Int).Mod(buff, big.NewInt(module))
	}

	return messageBig
}

func readCryptedMessage() []*big.Int {
	data, _ := ioutil.ReadFile("crypted_message.txt")

	var element string
	var cryptedMessageIntoBigInt []*big.Int

	for i := 1; i < len(data); i++ {
		if string(data[i]) != " " {
			element = element + string(data[i])
		} else {

			buff := new(big.Int)
			buff.SetString(element, 10)

			cryptedMessageIntoBigInt = append(cryptedMessageIntoBigInt, buff)

			element = ""
		}
	}
	return cryptedMessageIntoBigInt
}

func bigIntToText(data []*big.Int) string {

	var text string

	for i := 0; i <= len(data)-1; i++ {
		c := data[i].Int64()
		text = text + string(c)
	}
	return text
}

func stringToByte(text string) []byte {
	buff := []byte(text)

	return buff
}
