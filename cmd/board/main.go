package main

import (
	"fmt"
	"github.com/friendsofgo/board-kata/errors"
	"github.com/friendsofgo/board-kata/parser"
	"github.com/friendsofgo/board-kata/writter"
	"log"

	"github.com/friendsofgo/board-kata"
)

func main() {
	msg, err := board.ReadInput("data/input.csv")
	if errors.IsFileNotFound(err) {
		fmt.Println("could not open file :(")
		log.Fatal(err)
	}

	finalMessage := ""

	var partial string
	for _, message := range msg {
		partial, err = parser.Parse(message)
		if errors.IsEmptyMessage(err) {
			log.Println("Got an empty message :O")
			continue
		}
		finalMessage += partial
	}

	err = writter.Write(addHtmlHeaders(finalMessage), "./example.html")
	if errors.IsCouldNotWriteFile(err) {
		fmt.Println("could not write output html file :(")
		log.Fatal(err)
	}

	fmt.Println("done!")
}

func addHtmlHeaders(input string) (output string) {
	return "<html><head><meta charset='utf-8'/></head><body>"+input+"</body></html>"
}
