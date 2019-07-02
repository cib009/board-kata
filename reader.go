package board

import (
	"bufio"
	"encoding/csv"
	"github.com/friendsofgo/board-kata/errors"
	"io"
	"os"
)

// ReadInput read the file with the inputs
func ReadInput(path string) ([]string, error) {
	var messages []string

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return messages, errors.WrapFileNotFound(err, "Impossible open file")
	}
	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = ';'
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		messages = append(messages, line[0])
	}

	return messages, nil
}
