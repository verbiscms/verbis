package console

import (
	"bufio"
	"cms/core/bootstrap"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

func Process() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)
		argument1 := ""
		argument2 := ""

		if strings.Contains(command, "make") {

			if !strings.Contains(command, ":") {
				color.Red.Println("Command not found")
			} else {
				stringSlice := strings.Split(command, ":")
				argSlice := strings.Split(stringSlice[1], " ")

				if len(argSlice) == 1  {
					color.Red.Println("Pass two arguments to make. E.g make:controller filename")
				} else {
					command = stringSlice[0]
					argument1 = argSlice[0]
					argument2 = argSlice[1]
				}
			}
		}

		switch command {
			case "help":
				Options()
				break
			case "clear":
				print("\033[H\033[2J")
				break
			case "serve":
				bootstrap.Load()
				break
			case "make":
				make(argument1, argument2)
				break
			case "migrate":
				fmt.Println("in")
				//migrate.Migrate.Load()
				//migrate.Migrate.Fresh()
			case "exit":
				os.Exit(0)
				break
			default:
				color.Red.Println("Command not found")
		}
	}
}

// Process arguments passed to console
func processArgs(command string) {

}

// Serve application
func serve(option string) {

}

// Make a module
func make(module string, name string) {

	switch module {
	case "migration":
		//migrate.Migrate.Load()
		//migrate.Migrate.Up()
		break
	case "controller":
		fmt.Print("In make controller")
		break
	case "model":
		fmt.Print("In make model")
		break
	}
}