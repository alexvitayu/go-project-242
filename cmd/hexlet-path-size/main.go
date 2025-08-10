package main

import (
	goproject242 "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	fmt.Println("Hello from Hexlet!")
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	fmt.Println(goproject242.GetSize("/home/alex/hexlet-git"))
}
