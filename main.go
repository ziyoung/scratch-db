package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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
		switch text {
		case ".exit":
			os.Exit(0)
		default:
			fmt.Printf("Unrecognized command '%s'.\n", text)
		}
	}
}
