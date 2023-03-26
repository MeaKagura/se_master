package main

import "fmt"

func main() {
	for {
		var cmd string
		fmt.Println("Please enter the command:")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("command list:")
			fmt.Println("--quit")
			fmt.Println("--hello")
		} else if cmd == "quit" {
			break
		} else if cmd == "hello" {
			fmt.Println("welcome to golang world!")
		} else {
			fmt.Println("Command not found. Input \"help\" for further help.")
		}
	}
}
