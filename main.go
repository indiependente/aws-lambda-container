package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) > 1 && os.Args[1] == "--test" {
		n, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			return err
		}
		fibN, err := fibonacci(n)
		if err != nil {
			return err
		}
		fmt.Println(fibN)
	} else {
		lambda.Start(handler)
	}
	return nil
}
