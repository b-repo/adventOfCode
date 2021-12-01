package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("2021/f/02.log")
	if err != nil {
		panic(err.Error())
	}

	data := strings.Split(string(f), "\n")

	fmt.Println(data)
}
