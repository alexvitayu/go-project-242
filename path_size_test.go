package goproject242

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name      string
	path_file string
	expect    string
	err       error
}{
	{
		name:      "Path-file",
		path_file: "testdata/text.txt",
		expect:    "23B\ttestdata/text.txt",
		err:       nil,
	},
	{
		name:      "Path-directory",
		path_file: "testdata",
		expect:    "82B\ttestdata",
		err:       nil,
	},
	{
		name:      "Empty path",
		path_file: "",
		expect:    "",
		err:       errors.New("не указан путь"),
	},
	{
		name:      "Wrong path",
		path_file: "testdata/text.txttt",
		expect:    "",
		err:       errors.New("не удалось прочитать путь к файлу или директории"),
	},
}

func TestGetSize(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetSize(tc.path_file)
			if tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("ожидали %v, получили %v", tc.err.Error(), err.Error())

				}
			}
			// Сравнение с помощью булевого значения (стандартно)
			/*if got != tc.expect {
				t.Errorf("ожидали %v, получили %v", tc.expect, got)
			}*/

			// Сравнение с помощью библиотеки require
			//require.Equal(t, tc.expect, got, "полученный результат не совпадает с ожидаемым")

			// Сравнение с помощью сторонней библиотеки testify
			assert.Equal(t, tc.expect, got, "they should be equal")
		})
	}
}
