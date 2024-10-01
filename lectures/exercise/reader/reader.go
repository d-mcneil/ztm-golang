//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	nonBlankLines, numCommands := 0, 0

	r := bufio.NewReader(os.Stdin)

	for {
		input, inputErr := r.ReadString('\n')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		if n == "Q" || n == "q" {
			break
		}
		nonBlankLines++
		if n == "hello" || n == "bye" {
			fmt.Println("Hello/Bye")
		}
		commands := strings.Split(n, " ")
		for _, command := range commands {
			command = strings.TrimSpace(command)
			if command != "" {
				numCommands++
			}
		}
		if inputErr != nil {
			if inputErr == io.EOF {
				break
			} else {
				fmt.Println(inputErr)
			}
		}
	}

	fmt.Printf("Non-blank lines: %d\n", nonBlankLines)
	fmt.Printf("Commands: %d\n", numCommands)
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	numLines, numCommands := 0, 0
// 	for scanner.Scan() {
// 		if strings.ToUpper(scanner.Text()) == "Q" {
// 			break
// 		}
// 		text := strings.ToLower(strings.TrimSpace(scanner.Text()))
// 		switch text {
// 		case "hello":
// 			numCommands++
// 			println("Hello accepted")
// 		case "bye":
// 			numCommands++
// 			println("Bye accepted")
// 		}
// 		if text != "" {
// 			numLines++
// 		}
// 	}
// 	fmt.Println("Number of non-blank lines entered: ", numLines)
// 	fmt.Println("Number of commands entered: ", numCommands)
// }
