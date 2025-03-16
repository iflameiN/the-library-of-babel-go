package main

import (
	"container/list"
	"fmt"
	"sync"
)

const (
	Letters      = "abcdefghijklmnopqrstuvwxyz ,. " // 31 characters
	BookLength   = 4096                             //4B book
	BooksPerHex  = 1024                             // Book per hexagon
	HexCacheSize = 128                              // Elem in lru cache
)

// pseudo-random number generator
type Xorshift struct {
	state uint64
}

func (x *Xorshift) Next() uint64 {
	x.state ^= x.state << 13
	x.state ^= x.state >> 7
	x.state ^= x.state << 17
	return x.state
}

type Book struct {
	ID    string
	HexID string
	Seed  uint64
}

type Hexagon struct {
	ID    string
	Books []Book
}

type HexCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mu sync.Mutex
}

func NewHexCache(capacity int) *HexCache {
	return &HexCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *HexCache) Get(key string) (*Hexagon, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		return el.Value.(*Hexagon), true
	}

	return nil, false
}

func (c *HexCache) Put(key string, value *Hexagon) {
	c.mu.Lock()
	defer c.mu.Unlock()

	
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		el.Value = value
		return 
	}

	newEl := c.list.PushFront(value)
	c.cache[key] = newEl

	if c.list.Len() > c.capacity {	
		lastEl := c.list.Back()

		if lastEl != nil {
			delete(c.cache, lastEl.Value.(*Hexagon).ID)
			c.list.Remove(lastEl)
		}
	}
}

//Global cache
var hexCache = NewHexCache(HexCacheSize);

//Genereta content of book
func (b *Book) GenerateContent(buf []byte) []byte {
	rng := Xorshift{state: b.Seed}

	for i := 0; i < BookLength; i++ {
		buf[i] = Letters[rng.Next()%uint64(len(Letters))]
	}

	return buf[:BookLength]
}

//Generate hexagon
func GeneterateHexagon(hexNumber int) *Hexagon {
	hexID := fmt.Sprintf("HEX-%d", hexNumber)
	books := make([]Book, BooksPerHex)

	var seedBase uint64 = uint64(hexNumber) * uint64(BooksPerHex)

	for i := range books {
		seed := seedBase + uint64(i)

		books[i] = Book{
			ID: fmt.Sprintf("HEX-%d-BOOK-%d", hexNumber, i),
			HexID: hexID,
			Seed: seed,
		}
	}

	return &Hexagon{
		ID:    hexID,
		Books: books,
	}
}


func GetHexagon(hexNumber int) *Hexagon {
	hexID := fmt.Sprintf("HEX-%d", hexNumber)

	if hex, found := hexCache.Get(hexID); found {
		return hex
	}

	hex := GeneterateHexagon(hexNumber)
	hexCache.Put(hexID, hex)
	return hex;
}

func main() {
	//Get hex
	hex := GetHexagon(12345)
	fmt.Printf("Hexagon %s contains %d books\n", hex.ID, len(hex.Books))

	book := hex.Books[678]
	buf := make([]byte, BookLength)
	content := book.GenerateContent(buf)
	fmt.Printf("Book %s starts with %s\n", book.ID, string(content[:20]))
}