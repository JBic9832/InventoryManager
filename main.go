package main

import "log"

func main() {
	s := NewServer(":8000")

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
