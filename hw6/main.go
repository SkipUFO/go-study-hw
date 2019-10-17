package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skipufo/dd"
)

var source string
var destination string
var offset int64
var limit int64
var bufferSize int64

func init() {
	flag.StringVar(&source, "source", "", "file to read from")
	flag.StringVar(&destination, "destination", "", "file to write to")
	flag.Int64Var(&offset, "offset", 0, "offset in source file")
	flag.Int64Var(&limit, "limit", 0, "limit bytes to write to destination")
	flag.Int64Var(&bufferSize, "buffer", 4096, "buffer size for copy. default 4096")
}

func main() {

	flag.Parse()

	// Проверяем корректность флагов
	if source == "" {
		fmt.Println("Source file not set")
		os.Exit(1)
	}

	if destination == "" {
		fmt.Println("destination file not set")
		os.Exit(1)
	}

	if offset < 0 {
		fmt.Println("offset must be above zero")
		os.Exit(1)
	}

	if limit < 0 {
		fmt.Println("limit must be above zero")
		os.Exit(1)
	}

	if bufferSize < 0 {
		fmt.Println("buffer must be above zero")
		os.Exit(1)
	}

	if source == destination {
		fmt.Println("source and destination must be different")
		os.Exit(1)
	}

	if err := dd.Copy(source, destination, offset, limit, bufferSize); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
