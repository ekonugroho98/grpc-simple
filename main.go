package main

import (

	"fmt"
	"os"
	db "github.com/ekonugroho/grpc-simple/repository"
)


func main() {
	fmt.Println("Hello world")
	if err := db.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
