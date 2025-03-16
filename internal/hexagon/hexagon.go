package hexagon

const (
	BookLength  = 4096
	BooksPerHex = 1024
)

type Hexagon struct {
	ID    string
	Books []Book
}

func (h *Hexagon) GetID() string {
	return h.ID
}