package main

import "github.com/laix0n/satim/internal/api/grpc"

func main() {
	s := grpc.NewServer()
	s.Run()
}
