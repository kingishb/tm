package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// run runs a shell command
func run(sh string) {
	// break command into parts
	words := strings.Split(sh, " ")
	cmd := exec.Command(words[0], words[1:]...)

	// set stdin etc to caller terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func main() {
	fmt.Println("=> current sessions:")
	run("tmux list-session")
	menu := `
=> a    - attach
   new  - new session
   kill - kill a session
   exit - exit
`

	for {
		fmt.Println(menu)
		fmt.Printf("=> ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "a":
			fmt.Println("=> name: ")
			var n string
			fmt.Scanln(&n)
			run(fmt.Sprintf("tmux a -t %s", n))
		case "new":
			fmt.Printf("=> name: ")
			var n string
			fmt.Scanln(&n)
			run(fmt.Sprintf("tmux new -s %s", n))
		case "kill":
			fmt.Println("=> session name: ")
			fmt.Printf("=> ")
			var k string
			fmt.Scanln(&k)
			run(fmt.Sprintf("tmux kill-session -t %s", k))
		case "exit":
			fmt.Println("=> goodbye!")
			return
		default:
			fmt.Println("invalid choice")
		}
		fmt.Println("=> current sessions:")
		run("tmux list-session")
	}

}
