package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func append() {
	errRemove := os.Remove("data.ln")
	check(errRemove)
	f, err := os.OpenFile("data.ln", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()
	f.WriteString("[")
	for i := 0; i < 1000; i++ {
		if i == 999 {
			if _, err = f.WriteString(fmt.Sprintf("{\"id\":%d, \"descript\":\"Test-%d\"}", i+1, i+1)); err != nil {
				panic(err)
			}
			continue
		}
		if _, err = f.WriteString(fmt.Sprintf("{\"id\":%d, \"descript\":\"Test-%d\"},", i+1, i+1)); err != nil {
			panic(err)
		}
	}
	f.WriteString("]")
}
func main() {
	append()
	dat, err := ioutil.ReadFile("data.ln")
	check(err)
	var types []interface{}
	errUnmarshal := json.Unmarshal(dat, &types)
	check(errUnmarshal)
	fmt.Printf("%+v", types)
}
