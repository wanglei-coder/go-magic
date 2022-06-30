package main

import (
	"flag"
	"log"

	"github.com/wanglei-coder/go-magic"
)

var filename string

func init() {
	flag.StringVar(&filename, "filename", "", "Please enter the file path")
	flag.Parse()
}

func main() {
	m, err := magic.NewMagic(magic.MagicMimeType | magic.MagicSymLink | magic.MagicError)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	got, err := m.FromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(got)
}
