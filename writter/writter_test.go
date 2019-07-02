package writter_test

import (
	"github.com/friendsofgo/board-kata/writter"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct{ message, path, expectedErrorMessage string }{
		"correct message": {"Es un test", "/tmp/test.txt", ""},
		"incorrect message": {"Es un test", "/tmp", "Impossible to write file: open /tmp: is a directory"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := writter.Write(tc.message, tc.path)

			if err == nil && tc.expectedErrorMessage != "" {
				t.Fatalf(
					"Wrong writter error expected: %s, got: %s",
					tc.expectedErrorMessage,
					err,
				)
			} else if err != nil && err.Error() != tc.expectedErrorMessage {
				t.Fatalf(
					"Wrong writter error expected: %s, got: %s",
					tc.expectedErrorMessage,
					err.Error(),
				)
			}
		})
	}
}
