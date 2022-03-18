package main

import (
	"io/ioutil"
	"log"
)

func writeData(path string, content []byte) {

	err := ioutil.WriteFile("./data/Users.json", content, 0644)

	if err != nil {
		log.Fatal(err)
	}

}
