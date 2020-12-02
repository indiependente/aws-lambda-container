package main

import "errors"

var (
	ErrInvalidInput = errors.New("negative input not valid")
)

// Copyright (c) 2015 Project Nayuki
// All rights reserved. Contact Nayuki for licensing.
// https://www.nayuki.io/page/fast-fibonacci-algorithms
func fibonacci(n int64) (uint64, error) {
	if n < 0 {
		return 0, ErrInvalidInput
	}
	fibN, _ := fastFibonacci(uint64(n))
	return fibN, nil
}

func fastFibonacci(n uint64) (uint64, uint64) {
	if n == 0 {
		return 0, 1
	}
	a, b := fastFibonacci(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b
	if n%2 == 0 {
		return c, d
	}
	return d, c + d
}
