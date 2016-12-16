package cache

import (
	"encoding/json"

	redis "gopkg.in/redis.v5"
)

var (
	client *redis.Client
)

// Config redis client
func Config(addr string, poolSize int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		PoolSize: poolSize,
	})
	return nil
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func NewBookFromJSON(j string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(j), &book)
	return book, err
}
