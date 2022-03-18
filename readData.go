package main

import (
	"io/ioutil"
	"log"
)

func readData(path string) []byte {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return content

}
