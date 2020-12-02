package main

import (
	"errors"
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
	if len(os.Args) == 1 {
		lambda.Start(handler)
		return nil
	}
	switch os.Args[1] {
	case "--test":
		n, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			return err
		}
		fibN, err := fibonacci(n)
		if err != nil {
			return err
		}
		fmt.Println(fibN)
	case "--apigw":
		lambda.Start(apiGWHandler)
	default:
		return errors.New("unknown flag")
	}
	return nil
}
