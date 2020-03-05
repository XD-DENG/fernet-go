package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fernet/fernet-go"
)

const Usage = `Usage: fernet-extractts

fernet-extractts takes token as input, and prints the timestamp for the token`

func main() {
	log.SetFlags(0)
	log.SetPrefix("fernet: ")

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	bb := make([]byte, base64.URLEncoding.DecodedLen(len(b)))
	n, _ := base64.URLEncoding.Decode(bb, b)
	t, err := fernet.ExtractTimestamp(bb[:n])
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Stdout.Write(append([]byte(t.Format(time.UnixDate)), '\n'))
	if err != nil {
		log.Fatalln(err)
	}
}
