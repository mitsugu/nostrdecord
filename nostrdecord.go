package main

import (
	"os"
	"fmt"
	"log"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	txt := os.Args[1]
	pref,tmp, err := nip19.Decode(txt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pref, tmp)
}

