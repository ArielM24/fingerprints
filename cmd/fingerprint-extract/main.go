package main

import (
	"encoding/json"
	"os"

	"github.com/alevinval/fingerprints/src/extraction"
	"github.com/alevinval/fingerprints/src/helpers"
)

func main() {
	path := os.Args[1]
	_, m := helpers.LoadImage(path)
	result := extraction.DetectionResult(m)
	d, _ := json.Marshal(result)
	println(string(d))
}
