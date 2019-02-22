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

// todo
func prepareStatement(input string) {
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
