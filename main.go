package main

import (
	"context"
	"s0609-23/internal/demodb"
)

func main() {
	storage := demodb.New(nil)

	storage.GetUserByID(context.Background(), 100500)
}
