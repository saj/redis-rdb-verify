package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Status struct {
	Message string
	Code    int
}

var (
	StatusOK = Status{
		Message: "ok",
		Code:    0,
	}
	StatusChecksumMismatch = Status{
		Message: "checksum mismatch",
		Code:    10,
	}
	StatusChecksumMissing = Status{
		Message: "checksum missing",
		Code:    11,
	}
)

func init() {
	log.SetFlags(0)
}

func main() {
	path := "-"
	rdb := bufio.NewReader(os.Stdin)
	if len(os.Args) > 1 {
		path = os.Args[1]
		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("open rdb: %s", err)
		}
		defer f.Close()
		rdb = bufio.NewReader(f)
	}

	computed, recorded, err := check(rdb)
	if err != nil {
		log.Fatal(err)
	}

	status := StatusOK
	if !recorded.Equals(computed) {
		status = StatusChecksumMismatch
	}
	if recorded.IsZero() {
		status = StatusChecksumMissing
	}
	fmt.Printf("%x\t%x\t%s\t%s\n", computed, recorded, status.Message, path)
	os.Exit(status.Code)
}
