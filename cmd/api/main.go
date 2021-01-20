package main

import (
	"basic-api/pkg/server"
	"fmt"
)

func main() {
	s := server.New()
	if err := s.Run(); err != nil {
		fmt.Printf("server error: %s", err)
	}
}
