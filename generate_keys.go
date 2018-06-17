package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Keys struct {
	firstPrimeNumber  int
	secondPrimeNumber int
	module            int64
	openKey           int64
	closedKey         int64
	eiler             int
	yourRandomNumber  int
}

func getKeys() {
	keys := Keys{}

	keys.firstPrimeNumber, keys.secondPrimeNumber = generatePrimeNumbers()
	keys.eiler, keys.module, keys.closedKey, keys.openKey = generateKeys(keys.firstPrimeNumber, keys.secondPrimeNumber)
	fmt.Println("Your closed key is - " + "{ " + strconv.FormatInt(keys.closedKey, 10) + " , " + strconv.FormatInt(keys.module, 10) + " }")
	fmt.Println("Your open key is - " + "{ " + strconv.FormatInt(keys.openKey, 10) + " , " + strconv.FormatInt(keys.module, 10) + " }")
	fmt.Println("Please save this keys for further usage. And keep closed key as safe as you can for the better security.")
}

func sieveOfEratosthenes(maxNumber int) (primes []int) {
	b := make([]bool, maxNumber)
	for i := 2; i < maxNumber; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < maxNumber; k += i {
			b[k] = true
		}
	}
	return
}

func generatePrimeNumbers() (firstPrimeNumber, secondPrimeNumber int) {

	primes := sieveOfEratosthenes(500)

	firstPrimeNumber = primes[rand.Intn(len(primes)-1)]
	secondPrimeNumber = primes[rand.Intn(len(primes)-1)]

	return
}

func generateKeys(firstPrimeNumber, secondPrimeNumber int) (eiler int, module64, closedKey64, openKey64 int64) {

	generatePrimeNumbers()

	module := firstPrimeNumber * secondPrimeNumber
	module64 = int64(module)

	eiler = (firstPrimeNumber - 1) * (secondPrimeNumber - 1)

	closedKey_ := 0
	closedKey := randomInt(200, eiler)
	for closedKey_ != 1 {
		closedKey++
		closedKey_ = gcd(closedKey, eiler)
	}

	closedKey64 = int64(closedKey)

	openKey := 0
	openKey_ := 0

	for openKey_ != 1 {
		openKey++
		openKey_ = (openKey * closedKey) % eiler
	}
	openKey64 = int64(openKey)

	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
