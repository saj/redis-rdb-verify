package crc64

import (
	"encoding/binary"
	"hash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CRC-64-Jones", func() {
	var h hash.Hash64

	BeforeEach(func() {
		h = New()
	})

	Describe("write", func() {
		var (
			input  []byte
			output []byte
		)

		JustBeforeEach(func() {
			h.Write(input)
			output = h.Sum(nil)
		})

		// This is a check value from the upstream Redis test suite.
		Context("check value", func() {
			BeforeEach(func() {
				input = []byte("123456789")
			})

			It("hashes to a known value", func() {
				expected := make([]byte, 8)
				binary.LittleEndian.PutUint64(expected, 0xe9c6d914c4b8d9ca)
				Expect(output).To(Equal(expected))
			})
		})
	})
})
