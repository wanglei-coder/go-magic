# go-magic
go-magic is a golang library that wraps [libmagic](https://linux.die.net/man/3/libmagic)

## Requirements
go-magic needs `libmagic` to be installed.
- On Debian or Ubuntu: `apt-get install -y libmagic-dev`
- On RHEL, CentOS or Fedora: `yum install file-devel`
- On Mac OS X: `brew install libmagic`

## Installing
```go get -u github.com/wanglei-coder/go-magic```

## Using the library
```go
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

	got, err := m.FromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(got)
}
```
