package main

import (
	"fmt"
	"log"
	"net/mail"
	"os"
)

func main() {
	message, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", message)
}
