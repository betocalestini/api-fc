package main

import (
	"fmt"

	"github.com/betocalestini/api-fc/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.DBDriver)
}
