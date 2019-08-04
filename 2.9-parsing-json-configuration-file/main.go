package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func (c *configuration) String() string {
	return fmt.Sprintf("%t - %s", c.Enabled, c.Path)
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.String())
}
