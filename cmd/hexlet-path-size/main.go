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
		UseShortOptionHandling: true,
		Name:                   "hexlet-path-size",
		Usage:                  "print size of a file or directory;",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := "." //default value
			if cmd.Args().Len() > 0 {
				path = cmd.Args().Get(0)
			}
			h := cmd.Bool("human")
			size, err := goproject242.GetSize(path, h)
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Println(size)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
