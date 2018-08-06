package main

import (
	"io/ioutil"
	"encoding/json"
)

//TODO add file watcher and filename input
func main() {
	file,_ := ioutil.ReadFile("input.json")

	input := &PuzzleSetup{}

	json.Unmarshal([]byte(string(file)), &input)

	input.ParseOps()
	input.Solve()
}

