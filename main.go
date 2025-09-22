package main

import (
	"fmt"
	"myapp_go/config"
)

func main() {
	fmt.Println("hello world")
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	fmt.Printf("config.Get().Http.Port: %v\n", config.Get().Http.Port)
}
