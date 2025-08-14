package code

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if path == "" {
		return "", errors.New("не указан путь")
	}
	sum, err := GetSize(path, recursive, human, all)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return fmt.Sprintf("%v", FormatSize(sum, human)), nil

}

func GetSize(path string, recursive, human, all bool) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, errors.New("не удалось прочитать путь к файлу или директории")
	}
	// If path is a file but not a directory
	if !info.IsDir() {
		i, err := os.Lstat(path)
		if err != nil {
			return 0, errors.New("не удалось получить информацию о файле")
		}
		return i.Size(), nil
	}

	// If path is a directory not a file
	var sum int64
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, errors.New("не удалось прочитать директорию")
	}
	for i := 0; i < len(files); i++ {
		if !all && strings.HasPrefix(files[i].Name(), ".") {
			continue
		}
		if !files[i].IsDir() {
			fP := filepath.Join(path, files[i].Name())
			stat, err := os.Lstat(fP)
			if err != nil {
				return 0, errors.New("не удалось получить информацию о файле")
			}
			sum += stat.Size()
		} else if recursive && files[i].IsDir() {
			fP := filepath.Join(path, files[i].Name())
			s, err := GetSize(fP, recursive, human, all)
			if err != nil {
				return 0, fmt.Errorf("%w", err)
			}
			sum += s
		}
	}
	return sum, nil
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
			return fmt.Sprintf("%.1fMB", float64(size)/math.Pow(10, 6))
		case len(str) >= 10 && len(str) < 13:
			return fmt.Sprintf("%.1fGB", float64(size)/math.Pow(10, 9))
		case len(str) >= 13 && len(str) < 16:
			return fmt.Sprintf("%.1fTB", float64(size)/math.Pow(10, 12))
		case len(str) >= 16 && len(str) < 19:
			return fmt.Sprintf("%.1fPB", float64(size)/math.Pow(10, 15))
		default:
			return fmt.Sprintf("%.1fEB", float64(size)/math.Pow(10, 18))
		}
	}
	return fmt.Sprintf("%vB", size)
}
