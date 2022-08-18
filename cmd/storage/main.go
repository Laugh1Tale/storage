package main

import (
	"fmt"
	"github.com/laugh1tale/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("It's works!", st)
}
