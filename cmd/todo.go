package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/smothiki/protobuftuts/todo"
)

func add() {
	task := &todo.Task{
		Text: "siva",
		Done: true,
	}
	fmt.Println(task)
}
func list() {}

const db = "db.pb"

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("no enough flags")
		os.Exit(1)
	}
	cmd := flag.Arg(1)
	switch cmd {
	case "list":
		list()
	case "add":
		add()

	}
	fmt.Println("Todo app")
}
