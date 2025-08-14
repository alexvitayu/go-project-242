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
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories"},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var path string
			if cmd.Args().Len() > 0 {
				path = cmd.Args().Get(0)
			}
			h := cmd.Bool("human")
			a := cmd.Bool("all")
			r := cmd.Bool("recursive")
			size, err := goproject242.GetPathSize(path, r, h, a)
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
