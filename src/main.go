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

/**
 * 在main 函数返回之前清理并终止之前启动的goroutine
 */
func main()  {
	fmt.Println("hello world")
	search.Run("president");
}
