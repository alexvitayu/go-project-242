package goproject242

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func Help() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func GetSize(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", errors.New("не удалось прочитать путь к файлу или директории")
	}
	if !info.IsDir() {
		i, err := os.Lstat(path)
		if err != nil {
			return "", errors.New("не удалось получить информацию о файле")
		}
		return fmt.Sprintf("%vB\t%s", i.Size(), path), nil
	}

	var sum int64
	files, err := os.ReadDir(path)
	if err != nil {
		return "", errors.New("не удалось прочитать директорию")
	}
	for _, file := range files {
		if !file.IsDir() {
			fP := filepath.Join(path, file.Name())
			stat, err := os.Lstat(fP)
			if err != nil {
				return "", errors.New("не удалось получить информацию о файле")
			}
			sum += stat.Size()
		}
	}
	return fmt.Sprintf("%vB\t%s", sum, path), nil
}
