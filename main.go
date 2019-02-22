package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func doMetaCommand(input string) {
	switch input {
	case ".exit":
		os.Exit(0)
	default:
		fmt.Printf("Unrecognized command '%s'\n", input)
	}
}

// Statement is Query data
type Statement struct {
	typ string
}

// Execute query
func (s Statement) Execute() {
	switch s.typ {
	case "select":
		fmt.Println("This is where we would do an insert.")
	case "insert":
		fmt.Println("This is where we would do a select.")
	default:
		fmt.Printf("Unsupported Statement type %s\n", s.typ)
	}
}

func prepareStatement(input string) (*Statement, error) {
	statement := &Statement{}
	typ := ""
	if strings.HasPrefix(input, "select") {
		typ = "select"
	} else if strings.HasPrefix(input, "insert") {
		typ = "insert"
	}
	if typ != "" {
		statement.typ = typ
		return statement, nil
	}
	return nil, errors.New("Unrecognized command")
}

func main() {
	fmt.Println("Server is running.")
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "readString error %v\n", err)
			os.Exit(1)
		}
		text = strings.Replace(text, "\n", "", -1)
		if strings.HasPrefix(text, ".") {
			doMetaCommand(text)
			continue
		}
		statement, err := prepareStatement(text)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		statement.Execute()
		fmt.Println("Executed.")
	}
}
