package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var types []string
	input := []byte(`["one","two","many"]`)
	json.Unmarshal(input, &types)
	fmt.Printf("%+v", types)
}
