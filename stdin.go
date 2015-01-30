package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Println(`command:
	print <content> "print your words"
	plus <number> <number> "help you to plus two number"
	quit "quit"`)
}

func echo(words string) {
	fmt.Println(words)
}

func plus(one, two int) {
	all := one + two
	fmt.Println(one, "+", two, "is", all)
}

func main() {
	help()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)
		token := strings.Split(line, " ")
		count := len(token)
		if count > 0 {
			command := token[0]
			if command == "print" {
				if count > 1 {
					words := token[1]
					echo(words)
				} else {
					fmt.Println("No words to print.")
				}
			} else if command == "plus" {
				if count == 3 {
					one, err1 := strconv.Atoi(token[1])
					two, err2 := strconv.Atoi(token[2])
					if err1 != nil || err2 != nil {
						fmt.Println("please enter int numbers")
					} else {
						plus(one, two)
					}
				} else {
					fmt.Println("Only two int numbers work.")
				}
			} else if command == "quit" {
				break
			} else {
				fmt.Println("Unknown Command.")
			}
		}
	}
}
