package main

import
(
"math/rand"
"strconv"
"fmt"
"time"
"math/big"
"io/ioutil"
"flag"
)

type Keys struct{
  firstPrimeNumber int
  secondPrimeNumber int
  module int64
  openKey int64
  closedKey int64
  eiler int
  yourRandomNumber int
}

type Data struct{
  message []byte
  messageBig []*big.Int
  cryptedMessage []*big.Int
  decryptedMessage []*big.Int
  output []byte
  buff string
}

func main(){

    rand.Seed(time.Now().UnixNano())
    keys := Keys{}
    data := Data{}

    option := flag.String("option", "waiting", "Choose option 1. keys - for generating keys 2.crypt - for crypting message 3. decrypt - for decrypting message")
    openKey := flag.Int64("open", 1, "Enter your open key")
    closedKey := flag.Int64("closed", 1, "Enter your closed key")
    signature := flag.Int64("signature", 1, "Enter your signature")

    flag.Parse()

    switch *option {

    case "keys":
      keys.firstPrimeNumber, keys.secondPrimeNumber = generatePrimeNumbers()
      keys.eiler, keys.module, keys.closedKey, keys.openKey = generateKeys(keys.firstPrimeNumber, keys.secondPrimeNumber)
      fmt.Println("Your closed key is - " + "{ " + strconv.FormatInt(keys.closedKey, 10) + " , " + strconv.FormatInt(keys.module, 10) + " }")
      fmt.Println("Your open key is - " + "{ " + strconv.FormatInt(keys.openKey, 10) + " , " + strconv.FormatInt(keys.module, 10) + " }")
      fmt.Println("Please save this keys for further usage. And keep closed key as safe as you can for the better security.")

    case "crypt":
      data.message = readFile("message.txt")
      data.messageBig = byteIntoBigInt(data.message)
      data.cryptedMessage = cryptMessage(data.messageBig, *openKey, *signature)
      data.output = bigIntToByte(data.cryptedMessage)
      writeFile("crypted_message.txt", data.output)
      fmt.Println("Your message sucsefully crypted.")

    case "decrypt":
      data.decryptedMessage = readCryptedMessage()
      data.decryptedMessage = decryptMessage(data.decryptedMessage, *closedKey, *signature)
      data.buff = bigIntToText(data.decryptedMessage)
      data.output = stringToByte(data.buff)
      writeFile("decrypted_message.txt", data.output)
      fmt.Println("Your message sucsefully decrypted")

    case "waiting":
      fmt.Println("Please choose what you want to do.")

    }
}

func sieveOfEratosthenes(maxNumber int) (primes []int) {
    b := make([]bool, maxNumber)
    for i := 2; i < maxNumber; i++ {
        if b[i] == true { continue }
        primes = append(primes, i)
        for k := i * i; k < maxNumber; k += i {
            b[k] = true
        }
    }
    return
}

func generatePrimeNumbers() (firstPrimeNumber, secondPrimeNumber int){

  primes := sieveOfEratosthenes(500)

  firstPrimeNumber = primes [rand.Intn(len(primes)-1)]
  secondPrimeNumber = primes [rand.Intn(len(primes)-1)]

  return
}

func generateKeys(firstPrimeNumber, secondPrimeNumber int) ( eiler int, module64, closedKey64, openKey64 int64){

    generatePrimeNumbers()

    module := firstPrimeNumber * secondPrimeNumber
    module64 = int64(module)

    eiler = (firstPrimeNumber - 1) * (secondPrimeNumber - 1)

    closedKey_ := 0
    closedKey := randomInt(200, eiler)
    for closedKey_ != 1{
      closedKey ++
      closedKey_ = gcd(closedKey, eiler)
    }

    closedKey64 = int64(closedKey)

    openKey := 0
    openKey_ := 0

    for openKey_ != 1{
      openKey ++
      openKey_ = (openKey * closedKey) % eiler
    }
    openKey64 = int64(openKey)

    return
}

func gcd(a, b int) int {
  for b != 0{
    a, b = b, a % b
    }
    return a
}

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func readFile(fileName string) []byte{
  file, err := ioutil.ReadFile(fileName) // just pass the file name
        if err != nil {
            fmt.Print(err)
            }

  return file
}

func writeFile(filePath string, data []byte){
  ioutil.WriteFile(filePath, data, 0644)
}

func byteIntoBigInt( message []byte ) []*big.Int {
  j := 1

  var messageIntoBigInt []*big.Int

  //convert []byte message into the []*big.Int message
  for i := 0 ; i < len(message) ; i++ {

    buff := new(big.Int)
    buff.SetBytes(message[i : j])
    j++

    messageIntoBigInt = append(messageIntoBigInt, buff)
  }
  return messageIntoBigInt
}

func bigIntToByte(file []*big.Int) (outputByte []byte){
  output := fmt.Sprint(file)
  outputByte = []byte(output)
  return
}

func cryptMessage( messageBig []*big.Int, openKey, module int64 ) []*big.Int{



    for i := 0 ; i < len(messageBig) ; i++{

    //power
    buff := (new(big.Int).Exp(messageBig[i], big.NewInt(openKey), nil))
    //%
    messageBig [i]= new(big.Int).Mod(buff, big.NewInt(module))

    }

    for i := 0 ; i < len(messageBig) - 1 ; i++{
      buff := (new(big.Int).Add(messageBig[i + 1], messageBig[i]))
      messageBig [i + 1]= new(big.Int).Mod(buff, big.NewInt(module))
    }

    return messageBig
}

func decryptMessage( messageBig []*big.Int, closedKey, module int64 ) []*big.Int{

  for i := len(messageBig) - 1; i > 0; i --{
    buff := (new(big.Int).Sub(messageBig[i], messageBig[i - 1]))
    messageBig [i]= new(big.Int).Mod(buff, big.NewInt(module))
  }

  for i := 0 ; i < len(messageBig) ; i++{
  //power
  buff := (new(big.Int).Exp(messageBig[i], big.NewInt(closedKey), nil))
  //%
  messageBig [i]= new(big.Int).Mod(buff, big.NewInt(module))
  }


  return messageBig
}

func readCryptedMessage() []*big.Int{
  data, _ := ioutil.ReadFile("crypted_message.txt")

  var element string
  var cryptedMessageIntoBigInt []*big.Int

  for i := 1; i < len(data) ; i++{
    if string(data[i]) != " "{
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

func bigIntToText( data []*big.Int) string{

    var text string

    for i := 0; i <= len(data) - 1; i++{
      c := data[i].Int64()
      text = text + string(c)
    }
    return text
}

func stringToByte(text string) []byte{
  buff := []byte(text)

  return buff
}
