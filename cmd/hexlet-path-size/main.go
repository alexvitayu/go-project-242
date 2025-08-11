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

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
		path := os.Args[1]
		fmt.Println(goproject242.GetSize(path))
	}

}
