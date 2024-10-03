package main

import (
	"fmt"
	"log"
)

// underlying storage: memory, disk, s3 => looks like an interface is needed
// server: http, tcp
// client?

func main() {
	cfg := &Config{
		ListenAddr: ":3000",
		StoreProducerFunc: func() Storer { return NewMemoryStore() },
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
    s.Start()

    /*
	s.Store.Push([]byte("foobar"))
	// data, err := s.Store.Fetch(offset)
	data, err := s.Store.Fetch(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	offset, err := s.Store.Push([]byte("zoobaz"))
	data, err = s.Store.Fetch(offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println(offset)
    */
	fmt.Println("poop")
}
