package goproject242

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetSize(path string, human, all bool) (string, error) {
	if path == "" {
		return "", errors.New("не указан путь")
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", errors.New("не удалось прочитать путь к файлу или директории")
	}
	// If path is a file but not a directory
	if !info.IsDir() {
		i, err := os.Lstat(path)
		if err != nil {
			return "", errors.New("не удалось получить информацию о файле")
		}
		return fmt.Sprintf("%v\t%s", FormatSize(i.Size(), human), path), nil
	}

	// If path is a directory not a file
	var sum int64
	files, err := os.ReadDir(path)
	if err != nil {
		return "", errors.New("не удалось прочитать директорию")
	}
	for i := 0; i < len(files); i++ {
		if !all && strings.HasPrefix(files[i].Name(), ".") {
			continue
		}
		if !files[i].IsDir() {
			fP := filepath.Join(path, files[i].Name())
			stat, err := os.Lstat(fP)
			if err != nil {
				return "", errors.New("не удалось получить информацию о файле")
			}
			sum += stat.Size()
		}
	}
	return fmt.Sprintf("%v\t%s", FormatSize(sum, human), path), nil
}

func FormatSize(size int64, human bool) string {
	if human {
		str := strconv.Itoa(int(size))
		switch {
		case len(str) <= 3:
			return fmt.Sprintf("%vB", size)
		case len(str) > 3 && len(str) <= 6:
			return fmt.Sprintf("%.1fKB", float64(size)/1000)
		case len(str) >= 7 && len(str) < 10:
			return fmt.Sprintf("%.1fMB", float64(size)/1000000)
		case len(str) >= 10 && len(str) < 13:
			return fmt.Sprintf("%.1fGB", float64(size)/1000000000)
		case len(str) >= 13 && len(str) < 16:
			return fmt.Sprintf("%.1fTB", float64(size)/1000000000000)
		case len(str) >= 16 && len(str) < 19:
			return fmt.Sprintf("%.1fPB", float64(size)/1000000000000000)
		default:
			return fmt.Sprintf("%.1fEB", float64(size)/1000000000000000000)
		}
	}
	return fmt.Sprintf("%vB", size)
}
