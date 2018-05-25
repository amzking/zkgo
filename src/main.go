package main

import (
	"fmt"
	"log"
	"os"

	"zkgo/src/search"
)

func init() {
	fmt.Println("i'm a init.")
	log.SetOutput(os.Stdout)
}

func main()  {
	fmt.Println("hello world")
	search.Run("president");
}
