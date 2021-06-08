package main

import (
	"github.com/RainrainWu/hcpairing"
)

func main() {

	hcpairing.DBConn.Start()
	server := hcpairing.NewServer()
	server.Start()
}
