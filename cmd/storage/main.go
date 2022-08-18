package main

import (
	"fmt"
	"github.com/laugh1tale/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It's uploaded!", file)
}
