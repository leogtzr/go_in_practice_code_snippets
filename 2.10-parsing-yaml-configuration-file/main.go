package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(config.Get("path"))
	// fmt.Println(config.GetBool("enabled"))

	path, _ := config.Get("path")
	enabled, _ := config.GetBool("enabled")

	fmt.Println(path)
	fmt.Println(enabled)
}
