package main

import (
	_16 "2021/16"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	hexStringBuff, err := ioutil.ReadFile("data/16.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	h, err := hex.DecodeString(string(hexStringBuff))
	if err != nil {
		log.Fatal(err.Error())
	}

	b := _16.ParseBits(h)
	version, value, _, err := b.Decode()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Part 1 : %d\n", version)
	fmt.Printf("Part 2 : %d\n", value)
}
