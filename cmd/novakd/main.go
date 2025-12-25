package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().UTC().Format(time.RFC3339), "hello Novak")
		time.Sleep(5 * time.Second)
	}
}
