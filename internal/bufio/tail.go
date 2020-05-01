package bufio

type RotatingTail struct {
	l int

	b  []byte // working tail buffer
	nb int

	d []byte // discarded from b on last Append
}

func NewRotatingTail(length int) *RotatingTail {
	return &RotatingTail{
		l: length,
		b: make([]byte, length),
		d: make([]byte, length),
	}
}

func (t *RotatingTail) Bytes() []byte { return t.b[:t.nb] }

func (t *RotatingTail) Append(b []byte) (discarded, head []byte) {
	nb := len(b)
	n := t.l
	if n > nb {
		n = nb
	}
	is := nb - n
	var tail []byte
	head, tail = b[:is], b[is:]

	nd := n
	if nd > t.nb {
		nd = t.nb
	}
	if t.nb+n <= t.l {
		nd = 0
	}
	copy(t.d, t.b[:nd])
	discarded = t.d[:nd]

	if nd > 0 {
		copy(t.b, t.b[nd:])
		t.nb -= nd
	}
	if t.nb+n <= t.l {
		copy(t.b[t.nb:], tail)
	} else {
		copy(t.b, tail)
	}
	t.nb += n
	return
}
