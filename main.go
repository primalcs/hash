package main

import (
	"log"

	"github.com/primalcs/hash/hash"
)

func main() {
	res := hash.Hash("asdqsda", "asdwqe")
	log.Println(res)
}
