package main

import (
	"github.com/lai0xn/satim-dolphin/listener"
	"github.com/lai0xn/satim-dolphin/store"
)

func main() {
	client := store.NewRedisStore()
	listener.Listen(client)
}
