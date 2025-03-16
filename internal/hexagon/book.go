package hexagon

import "vavilon-library/internal/rng"

const (
	Letters = "abcdefghijklmnopqrstuvwxyz ,. "
)

type Book struct {
	ID    string
	HexID string
	Seed  uint64
}

func (b *Book) GenerateContent(buf []byte) []byte {
	rng := rng.NewXorshift(b.Seed)
	for i := 0; i < BookLength; i++ {
		buf[i] = Letters[rng.Next()%uint64(len(Letters))]
	}
	return buf[:BookLength]
}