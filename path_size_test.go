package code

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name      string
	path_file string
	expect    string
	err       error
	human     bool
	all       bool
	recursive bool
}{
	{
		name:      "Path-file",
		path_file: "testdata/text.txt",
		expect:    "23B",
		err:       nil,
		human:     true,
		all:       true,
		recursive: false,
	},
	{
		name:      "Path-directory",
		path_file: "testdata",
		expect:    "9153B",
		err:       nil,
		human:     false,
		all:       true,
		recursive: false,
	},
	{
		name:      "Empty path",
		path_file: "",
		expect:    "",
		err:       errors.New("не указан путь"),
		human:     true,
		all:       true,
		recursive: false,
	},
	{
		name:      "Wrong path",
		path_file: "testdata/text.txttt",
		expect:    "",
		err:       errors.New("не удалось прочитать путь к файлу или директории"),
		human:     true,
		all:       true,
		recursive: false,
	},
	{
		name:      "Without hidden files",
		path_file: "testdata",
		expect:    "8939B",
		err:       nil,
		human:     false,
		all:       false,
		recursive: false,
	},
	{
		name:      "Recursive",
		path_file: "testdata",
		expect:    "6947251B",
		err:       nil,
		human:     false,
		all:       true,
		recursive: true,
	},
}

func TestGetPathSize(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetPathSize(tc.path_file, tc.recursive, tc.human, tc.all)
			if tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("ожидали %v, получили %v", tc.err.Error(), err.Error())

				}
			}
			// Сравнение с помощью булевого значения (стандартно)
			if got != tc.expect {
				t.Errorf("ожидали %v, получили %v", tc.expect, got)
			}

			// Сравнение с помощью библиотеки require
			//require.Equal(t, tc.expect, got, "полученный результат не совпадает с ожидаемым")

			// Сравнение с помощью сторонней библиотеки testify
			assert.Equal(t, tc.expect, got, "they should be equal")
		})
	}
}

var testCases2 = []struct {
	name   string
	size   []int64
	human  bool
	expect []string
}{
	{
		name:   "nonHumanReadable",
		size:   []int64{123, 10000, 500000},
		human:  false,
		expect: []string{"123B", "10000B", "500000B"},
	},
	{
		name:   "HumanReadable",
		size:   []int64{123, 10000, 999800, 11000000820},
		human:  true,
		expect: []string{"123B", "10.0KB", "999.8KB", "11.0GB"},
	},
}

func TestFormatSize(t *testing.T) {
	for _, tc := range testCases2 {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for _, s := range tc.size {
				size := FormatSize(s, tc.human)
				got = append(got, size)
			}
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("ожидали %v, получили %v", tc.expect, got)
			}
		})
	}
}
