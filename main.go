package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

//go:embed emojis.json
var raw []byte

func main() {
	var emojis []string
	err := json.Unmarshal(raw, &emojis)
	if err != nil || len(emojis) == 0 {
		fmt.Print("⚠")
		os.Exit(0)
	}

	n := rand.Int() % len(emojis)
	fmt.Print(emojis[n])
}
