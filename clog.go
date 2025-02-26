package clog

import "fmt"

const (
	RESET = "\033[0m"
	BOLD  = "\033[1m"
	BLACK = "\033[30m"
	RED   = "\033[31m"
)

func Bold(msg string) {
	fmt.Println(fmt.Sprintf("%s%s%s", BOLD, msg, RESET))
}

func Black(msg string) {
	fmt.Println(fmt.Sprintf("%s%s%s", BLACK, msg, RESET))
}

func Red(msg string) {
	fmt.Println(fmt.Sprintf("%s%s%s", RED, msg, RESET))
}
