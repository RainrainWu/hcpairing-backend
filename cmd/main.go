package main

import (
	"fmt"

	"github.com/RainrainWu/hcpairing"
)

func main() {

	hcpairing.DBConn.Start()
	hcpairing.DBConn.SetupTags()
	// hcpairing.DBConn.AppendRecord("california", []string{"Cough", "Itchy Skin"})
	for _, item := range hcpairing.DBConn.GetRecords() {
		fmt.Println(item)
	}

	server := hcpairing.NewServer()
	server.Start()
}
