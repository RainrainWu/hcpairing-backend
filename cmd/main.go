package main

import (
	"github.com/RainrainWu/hcpairing"
)

func main() {

	server := hcpairing.NewServer()
	server.Start()
}