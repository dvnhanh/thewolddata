package main

import (
	"log"

	"github.com/dvnhanh/thewolddata/cmd/cmdTheworlddata"
)

func main() {
	if err := cmdTheworlddata.Execute(); err != nil {
		log.Fatalf("Some errors wile execute: %s", err.Error())
	}
}
