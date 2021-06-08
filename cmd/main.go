package main

import (
	"fmt"

	"github.com/RainrainWu/hcpairing"
)

func main() {

	hcpairing.DBConn.Start()
	// hcpairing.DBConn.AppendRecord("california", []string{"Cough", "Itchy Skin"})
	for _, item := range hcpairing.DBConn.GetRecordsByState("california") {
		fmt.Println(item)
	}

	server := hcpairing.NewServer()
	server.Start()
}
