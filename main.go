package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// run runs a shell command
func run(sh string) error {
	// break command into parts
	words := strings.Split(sh, " ")
	cmd := exec.Command(words[0], words[1:]...)

	// set stdin etc to caller terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("=> current sessions:")
	err := run("tmux list-session")
	if err != nil {
		log.Fatal(err)
	}
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
			fmt.Printf("=> name: ")
			var n string
			fmt.Scanln(&n)
			err = run(fmt.Sprintf("tmux a -t %s", n))
			if err != nil {
				log.Fatal(err)
			}
		case "new":
			fmt.Printf("=> name: ")
			var n string
			fmt.Scanln(&n)
			err = run(fmt.Sprintf("tmux new -s %s", n))
			if err != nil {
				log.Fatal(err)
			}
		case "kill":
			fmt.Println("=> session name: ")
			fmt.Printf("=> ")
			var k string
			fmt.Scanln(&k)
			err = run(fmt.Sprintf("tmux kill-session -t %s", input))
			if err != nil {
				log.Fatal(err)
			}
		case "exit":
			fmt.Println("=> goodbye!")
			return
		default:
			fmt.Println("invalid choice")
		}
	}

}
