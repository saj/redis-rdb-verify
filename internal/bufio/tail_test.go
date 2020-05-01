package bufio

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RotatingTail", func() {
	var t *RotatingTail

	BeforeEach(func() {
		t = NewRotatingTail(8)
	})

	Context("with an empty buffer", func() {
		Context("after appending zero bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{})
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer is empty", func() {
				Expect(len(b)).To(Equal(0))
			})
		})

		Context("after appending three bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{
					0x00, 0x11, 0x22,
				})
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains all appended bytes", func() {
				Expect(len(b)).To(Equal(3))
				Expect(b[0]).To(Equal(byte(0x00)))
				Expect(b[1]).To(Equal(byte(0x11)))
				Expect(b[2]).To(Equal(byte(0x22)))
			})
		})

		Context("after appending eight bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 8)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains all appended bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(0)))
				Expect(b[1]).To(Equal(byte(1)))
				Expect(b[2]).To(Equal(byte(2)))
				Expect(b[3]).To(Equal(byte(3)))
				Expect(b[4]).To(Equal(byte(4)))
				Expect(b[5]).To(Equal(byte(5)))
				Expect(b[6]).To(Equal(byte(6)))
				Expect(b[7]).To(Equal(byte(7)))
			})
		})

		Context("after appending twenty bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 20)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head contains first twelve bytes", func() {
				Expect(len(head)).To(Equal(12))
				Expect(head[0]).To(Equal(byte(0)))
				Expect(head[11]).To(Equal(byte(11)))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(12)))
				Expect(b[1]).To(Equal(byte(13)))
				Expect(b[2]).To(Equal(byte(14)))
				Expect(b[3]).To(Equal(byte(15)))
				Expect(b[4]).To(Equal(byte(16)))
				Expect(b[5]).To(Equal(byte(17)))
				Expect(b[6]).To(Equal(byte(18)))
				Expect(b[7]).To(Equal(byte(19)))
			})
		})
	})

	Context("with a half-filled buffer", func() {
		BeforeEach(func() {
			a := make([]byte, 4)
			for i := 0; i < len(a); i++ {
				a[i] = byte(0xf0 + i)
			}
			t.Append(a)
		})

		Context("after appending zero bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{})
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer is unchanged", func() {
				Expect(len(b)).To(Equal(4))
				Expect(b[0]).To(Equal(byte(0xf0)))
				Expect(b[1]).To(Equal(byte(0xf1)))
				Expect(b[2]).To(Equal(byte(0xf2)))
				Expect(b[3]).To(Equal(byte(0xf3)))
			})
		})

		Context("after appending three bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{
					0x00, 0x11, 0x22,
				})
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains all appended bytes", func() {
				Expect(len(b)).To(Equal(7))
				Expect(b[0]).To(Equal(byte(0xf0)))
				Expect(b[1]).To(Equal(byte(0xf1)))
				Expect(b[2]).To(Equal(byte(0xf2)))
				Expect(b[3]).To(Equal(byte(0xf3)))
				Expect(b[4]).To(Equal(byte(0x00)))
				Expect(b[5]).To(Equal(byte(0x11)))
				Expect(b[6]).To(Equal(byte(0x22)))
			})
		})

		Context("after appending eight bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 8)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("four bytes were discarded", func() {
				Expect(len(discarded)).To(Equal(4))
				Expect(discarded[0]).To(Equal(byte(0xf0)))
				Expect(discarded[1]).To(Equal(byte(0xf1)))
				Expect(discarded[2]).To(Equal(byte(0xf2)))
				Expect(discarded[3]).To(Equal(byte(0xf3)))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(0)))
				Expect(b[1]).To(Equal(byte(1)))
				Expect(b[2]).To(Equal(byte(2)))
				Expect(b[3]).To(Equal(byte(3)))
				Expect(b[4]).To(Equal(byte(4)))
				Expect(b[5]).To(Equal(byte(5)))
				Expect(b[6]).To(Equal(byte(6)))
				Expect(b[7]).To(Equal(byte(7)))
			})
		})

		Context("after appending twenty bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 20)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("four bytes were discarded", func() {
				Expect(len(discarded)).To(Equal(4))
				Expect(discarded[0]).To(Equal(byte(0xf0)))
				Expect(discarded[1]).To(Equal(byte(0xf1)))
				Expect(discarded[2]).To(Equal(byte(0xf2)))
				Expect(discarded[3]).To(Equal(byte(0xf3)))
			})
			Specify("head contains first twelve bytes", func() {
				Expect(len(head)).To(Equal(12))
				Expect(head[0]).To(Equal(byte(0)))
				Expect(head[11]).To(Equal(byte(11)))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(12)))
				Expect(b[1]).To(Equal(byte(13)))
				Expect(b[2]).To(Equal(byte(14)))
				Expect(b[3]).To(Equal(byte(15)))
				Expect(b[4]).To(Equal(byte(16)))
				Expect(b[5]).To(Equal(byte(17)))
				Expect(b[6]).To(Equal(byte(18)))
				Expect(b[7]).To(Equal(byte(19)))
			})
		})
	})

	Context("with a full buffer", func() {
		BeforeEach(func() {
			a := make([]byte, 8)
			for i := 0; i < len(a); i++ {
				a[i] = byte(0xf0 + i)
			}
			t.Append(a)
		})

		Context("after appending zero bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{})
				b = t.Bytes()
			})

			Specify("nothing was discarded", func() {
				Expect(len(discarded)).To(Equal(0))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer is unchanged", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(0xf0)))
				Expect(b[1]).To(Equal(byte(0xf1)))
				Expect(b[2]).To(Equal(byte(0xf2)))
				Expect(b[3]).To(Equal(byte(0xf3)))
				Expect(b[4]).To(Equal(byte(0xf4)))
				Expect(b[5]).To(Equal(byte(0xf5)))
				Expect(b[6]).To(Equal(byte(0xf6)))
				Expect(b[7]).To(Equal(byte(0xf7)))
			})
		})

		Context("after appending three bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				discarded, head = t.Append([]byte{
					0x00, 0x11, 0x22,
				})
				b = t.Bytes()
			})

			Specify("three bytes were discarded", func() {
				Expect(len(discarded)).To(Equal(3))
				Expect(discarded[0]).To(Equal(byte(0xf0)))
				Expect(discarded[1]).To(Equal(byte(0xf1)))
				Expect(discarded[2]).To(Equal(byte(0xf2)))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(0xf3)))
				Expect(b[1]).To(Equal(byte(0xf4)))
				Expect(b[2]).To(Equal(byte(0xf5)))
				Expect(b[3]).To(Equal(byte(0xf6)))
				Expect(b[4]).To(Equal(byte(0xf7)))
				Expect(b[5]).To(Equal(byte(0x00)))
				Expect(b[6]).To(Equal(byte(0x11)))
				Expect(b[7]).To(Equal(byte(0x22)))
			})
		})

		Context("after appending eight bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 8)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("eight bytes were discarded", func() {
				Expect(len(discarded)).To(Equal(8))
				Expect(discarded[0]).To(Equal(byte(0xf0)))
				Expect(discarded[1]).To(Equal(byte(0xf1)))
				Expect(discarded[2]).To(Equal(byte(0xf2)))
				Expect(discarded[3]).To(Equal(byte(0xf3)))
				Expect(discarded[4]).To(Equal(byte(0xf4)))
				Expect(discarded[5]).To(Equal(byte(0xf5)))
				Expect(discarded[6]).To(Equal(byte(0xf6)))
				Expect(discarded[7]).To(Equal(byte(0xf7)))
			})
			Specify("head is empty", func() {
				Expect(len(head)).To(Equal(0))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(0)))
				Expect(b[1]).To(Equal(byte(1)))
				Expect(b[2]).To(Equal(byte(2)))
				Expect(b[3]).To(Equal(byte(3)))
				Expect(b[4]).To(Equal(byte(4)))
				Expect(b[5]).To(Equal(byte(5)))
				Expect(b[6]).To(Equal(byte(6)))
				Expect(b[7]).To(Equal(byte(7)))
			})
		})

		Context("after appending twenty bytes", func() {
			var discarded, head, b []byte

			BeforeEach(func() {
				a := make([]byte, 20)
				for i := 0; i < len(a); i++ {
					a[i] = byte(i)
				}
				discarded, head = t.Append(a)
				b = t.Bytes()
			})

			Specify("eight bytes were discarded", func() {
				Expect(len(discarded)).To(Equal(8))
				Expect(discarded[0]).To(Equal(byte(0xf0)))
				Expect(discarded[1]).To(Equal(byte(0xf1)))
				Expect(discarded[2]).To(Equal(byte(0xf2)))
				Expect(discarded[3]).To(Equal(byte(0xf3)))
				Expect(discarded[4]).To(Equal(byte(0xf4)))
				Expect(discarded[5]).To(Equal(byte(0xf5)))
				Expect(discarded[6]).To(Equal(byte(0xf6)))
				Expect(discarded[7]).To(Equal(byte(0xf7)))
			})
			Specify("head contains first twelve bytes", func() {
				Expect(len(head)).To(Equal(12))
				Expect(head[0]).To(Equal(byte(0)))
				Expect(head[11]).To(Equal(byte(11)))
			})
			Specify("tail buffer contains last eight bytes", func() {
				Expect(len(b)).To(Equal(8))
				Expect(b[0]).To(Equal(byte(12)))
				Expect(b[1]).To(Equal(byte(13)))
				Expect(b[2]).To(Equal(byte(14)))
				Expect(b[3]).To(Equal(byte(15)))
				Expect(b[4]).To(Equal(byte(16)))
				Expect(b[5]).To(Equal(byte(17)))
				Expect(b[6]).To(Equal(byte(18)))
				Expect(b[7]).To(Equal(byte(19)))
			})
		})
	})
})
