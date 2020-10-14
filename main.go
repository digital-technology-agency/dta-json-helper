package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	var types []string
	json.Unmarshal(dat, &types)
	fmt.Printf("%+v", types)
}
