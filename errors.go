package main

import (
	"errors"
	"fmt"
	"os"
)

type PleaseError struct {
	Err      error
	ExitCode int
}

var (
	fewArgsErrMsg       = errors.New("too few args")
	invalidSyntaxErrMsg = errors.New("the argument is null or has syntax errors")
)

func FatalError(err PleaseError) {
	fmt.Println("please: error:", err.Err)
	os.Exit(err.ExitCode)
}
