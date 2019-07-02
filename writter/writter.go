package writter

import (
	"github.com/friendsofgo/board-kata/errors"
	"io/ioutil"
)

// Parse fomating the text into a valid output text
func Write(msg string, path string) (err error) {
	err = ioutil.WriteFile(path, []byte(msg), 0644)

	if err != nil {
		return errors.WrapCouldNotWriteFile(err, "Impossible to write file")
	}

	return nil
}
