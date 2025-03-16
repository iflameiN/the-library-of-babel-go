package hexagon

import (
	"fmt"
	"vavilon-library/internal/cache"
)

var hexCache cache.Cache

func init() {
    hexCache = cache.NewHexCache(128)
}

func GenerateHexagon(hexNumber int) *Hexagon {
    hexID := fmt.Sprintf("HEX-%d", hexNumber)
    books := make([]Book, BooksPerHex)

    var seedBase uint64 = uint64(hexNumber) * uint64(BooksPerHex)
    for i := range books {
        seed := seedBase + uint64(i)
        books[i] = Book{
            ID:    fmt.Sprintf("HEX-%d-BOOK-%d", hexNumber, i),
            HexID: hexID,
            Seed:  seed,
        }
    }

    return &Hexagon{
        ID:    hexID,
        Books: books,
    }
}

func GetHexagon(hexNumber int) *Hexagon {
    hexID := fmt.Sprintf("HEX-%d", hexNumber)
    if cachedHex, found := hexCache.Get(hexID); found {
        return cachedHex.(*Hexagon)
    }

    hex := GenerateHexagon(hexNumber)
    hexCache.Put(hexID, hex)
    return hex
}