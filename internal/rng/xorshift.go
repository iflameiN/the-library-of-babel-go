package rng

type Xorshift struct {
	state uint64
}

func NewXorshift(seed uint64) *Xorshift {
	return &Xorshift{state: seed}
}

func (x *Xorshift) Next() uint64 {
	x.state ^= x.state << 13
	x.state ^= x.state >> 7
	x.state ^= x.state << 17
	return x.state
}