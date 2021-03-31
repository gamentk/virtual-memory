package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	page  []string
	stack []string
	hit   int
	fault int
)

func initialized() {
	stack = make([]string, 3)
	page = make([]string, 3)
	hit = 0
	fault = 0

	stack[0] = "null"
	stack[1] = "null"
	stack[2] = "null"
	page[0] = "null"
	page[1] = "null"
	page[2] = "null"
}

func showProcess() {
	fmt.Printf("\n\n| %s |\n| %s |\n| %s |\n", page[0], page[1], page[2])
	fmt.Printf("page-fault = %d\n", fault)
	fmt.Printf("hit = %d\n\n", hit)
	fmt.Printf("Least -> | %s | %s | %s | <- Most\n", stack[0], stack[1], stack[2])
	fmt.Printf("> ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func commandCreate(p string) {
	if page[0] == "null" && page[1] == "null" && page[2] == "null" {
		page[0] = p
		stack[0] = p
		fault++
	} else if page[0] != "null" && page[1] == "null" && page[2] == "null" {
		page[1] = p
		stack[1] = p
		fault++
	} else if page[0] != "null" && page[1] != "null" && page[2] == "null" {
		page[2] = p
		stack[2] = p
		fault++
	} else {
		if p != page[0] && p != page[1] && p != page[2] { // Page Fault
			for i := range page {
				if page[i] == stack[0] {
					page[i] = p
				}
			}
			stack[0] = stack[1]
			stack[1] = stack[2]
			stack[2] = p

			fault++
		} else { // Hit
			if p == stack[2] {
				hit++
			} else if p == stack[1] {
				stack[1] = stack[2]
				stack[2] = p
				hit++
			} else {
				stack[0] = stack[1]
				stack[1] = stack[2]
				stack[2] = p
				hit++
			}
		}
	}
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()

		switch command {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			commandCreate(command)
		case "exit", "q", "quit", "ex":
			return
		default:
			fmt.Printf("\nInvalid command.\n")
		}
	}
}
