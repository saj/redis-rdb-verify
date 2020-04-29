package bufio

type RotatingTail struct {
	l int
	b []byte // working tail buffer
	d []byte // discarded from b on last Append
}

func NewRotatingTail(length int) *RotatingTail {
	return &RotatingTail{l: length}
}

func (t *RotatingTail) Bytes() []byte { return t.b }

func (t *RotatingTail) Append(b []byte) (discarded, head []byte) {
	nb := len(b)
	nd := t.l
	if nd > nb {
		nd = nb
	}
	is := nb - nd
	var tail []byte
	head, tail = b[:is], b[is:nb]

	if nd == t.l {
		t.d = copyLazyAlloc(t.d, t.b, t.l)
		t.b = copyLazyAlloc(t.b, tail, t.l)
		discarded = t.d
		return
	}

	t.d = copyLazyAlloc(t.d, t.b[:nd], t.l)
	for i := 0; i < nd; i++ {
		t.b[i] = t.b[i+1]
	}
	t.b = copyLazyAlloc(t.b[nd:], tail, t.l)
	discarded = t.d[:nd]
	return
}

func copyLazyAlloc(dst, src []byte, l int) []byte {
	if len(src) == 0 {
		return dst
	}
	if dst == nil {
		dst = make([]byte, l)
	}
	copy(dst, src)
	return dst
}
