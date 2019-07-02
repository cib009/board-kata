package writter

import (
	"io/ioutil"
)

// Parse fomating the text into a valid output text
func Write(msg string, path string) {
	err := ioutil.WriteFile(path, []byte(msg), 0644)

	if err != nil {
		//Do SOmething
	}
}
