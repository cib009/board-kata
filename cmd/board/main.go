package main

import (
	"fmt"
	"github.com/friendsofgo/board-kata/parser"
	"github.com/friendsofgo/board-kata/writter"
	"log"

	"github.com/friendsofgo/board-kata"
)

func main() {
	msg, err := board.ReadInput("data/input.csv")
	if err != nil {
		log.Fatal(err)
	}

	finalMessage := ""

	for _, message :=range msg {
		finalMessage += parser.Parse(message)
	}

	writter.Write(finalMessage, "/Users/carles.calsina/go/src/github.com/board-kata/example.html")

	fmt.Println("done!")
}
