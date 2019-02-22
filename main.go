package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func doMetaCommand(input string) error {
	switch input {
	case ".exit":
		os.Exit(0)
		// return nil
	default:
		return fmt.Errorf("Unrecognized command '%s'", input)
	}
	return nil
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
		fmt.Fprintf(os.Stderr, "Unsupported Statement type %s\n", s.typ)
	}
}

// todo
func prepareStatement(input string) (Statement, error) {
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
			// excute meta command
			if err := doMetaCommand(text); err != nil {
				fmt.Fprint(os.Stderr, err)
			}
			break
		}
		// prepareStatement
	}
}
