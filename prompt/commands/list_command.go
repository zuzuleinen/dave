package commands

import "fmt"

func List() {
	fmt.Println("The list of commands are:")
	fmt.Println()
	fmt.Println("\tinstall:        Install the SQLite database.")
	fmt.Println("\tremind:         Create a new reminder.")
	fmt.Println("\treminders:      List all pending reminders.")
	fmt.Println("\tcredential:     Add a new credential record.")
	fmt.Println("\tcredentials:    List all credentials.")
	fmt.Println()
}
