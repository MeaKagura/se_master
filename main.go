package main

import (
	"fmt"
	"os"

	"github.com/MeaKagura/se_master/linklist"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN    = 1024
	CMD_NUM     = 10
)

func main() {
	for {
		cmd := make([]byte, CMD_MAX_LEN)
		fmt.Print(">>> Please input a command: ")
		fmt.Scanln(&cmd)
		p := linklist.FindCmd(head, string(cmd))
		if p == nil {
			fmt.Println("Command not fount, input \"help\" for further help.")
			continue
		}
		fmt.Printf("%s \n", p.Desc)
		if p.Cmd == "help" {
			help()
		}
		if p.Handler != nil {
			p.Handler()
		}
	}
}

var head *linklist.DataNode = &linklist.DataNode{
	Cmd:     "help",
	Desc:    "get help for commands.",
	Handler: nil,
	Next: &linklist.DataNode{
		Cmd:     "version",
		Desc:    "version of menu is 1.0.",
		Handler: nil,
		Next: &linklist.DataNode{
			Cmd:     "hello",
			Desc:    "hello, world!",
			Handler: nil,
			Next: &linklist.DataNode{
				Cmd:     "quit",
				Desc:    "exiting the program",
				Handler: quit,
				Next:    nil,
			},
		},
	},
}

var help = func() int {
	linklist.ShowAllCmd(head)
	return 0
}

var quit = func() {
	fmt.Println("Quit menu.")
	os.Exit(0)
}
