package main

import (
	"fmt"
	"io"

	ibufio "github.com/saj/redis-rdb-verify/internal/bufio"
	"github.com/saj/redis-rdb-verify/internal/crc64"
)

type RDBMagic [9]byte

var (
	RDBVersion9 = RDBMagic{0x52, 0x45, 0x44, 0x49, 0x53, 0x30, 0x30, 0x30, 0x39}
)

const sumLength = 8

type Sum []byte

func (s Sum) Equals(other Sum) bool {
	if len(s) != len(other) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != other[i] {
			return false
		}
	}
	return true
}

func (s Sum) IsZero() bool {
	for _, b := range s {
		if b != 0 {
			return false
		}
	}
	return true
}

func check(rdb io.Reader) (computed, recorded Sum, err error) {
	h := crc64.New()

	var magic RDBMagic
	_, err = io.ReadFull(rdb, magic[:])
	if err != nil {
		err = fmt.Errorf("read rdb magic: %s", err)
		return
	}
	switch magic {
	case RDBVersion9:
	default:
		err = fmt.Errorf("unknown rdb magic: %# x", magic)
		return
	}
	h.Write(magic[:])

	// The recorded checksum is found at the very end of the stream.
	var (
		t         = ibufio.NewRotatingTail(sumLength)
		b         = make([]byte, 4096)
		n         int
		readError error
	)
	for readError == nil {
		n, readError = rdb.Read(b)
		discarded, head := t.Append(b[:n])
		h.Write(discarded)
		h.Write(head)
	}
	computed = h.Sum(nil)

	if readError != io.EOF {
		err = fmt.Errorf("read: %s", readError)
		return
	}

	recorded = t.Bytes()
	return
}
