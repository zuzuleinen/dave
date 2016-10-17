package commands

import (
	"database/sql"
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/zuzuleinen/dave/credentials"
	"github.com/fatih/color"
)

func Credential(db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("User:")
	username := readFromPrompt(r)

	fmt.Print("Pass:")
	password := readFromPrompt(r)

	fmt.Print("Website:")
	website := readFromPrompt(r)

	fmt.Print("Notes:")
	notes := readFromPrompt(r)

	color.Green("Credentials succesfully added.")

	dbItems := []credentials.Credential{{username, password, website, notes}}
	credentials.Save(db, dbItems)
}

func ListCredentials(db *sql.DB) {
	items := credentials.Read(db)
	fmt.Println("Username | Password | Website | Notes\n")
	for _, item := range items {
		fmt.Println(item.Username, "|", item.Password, "|", item.Website, "|", item.Notes)
	}
}

func readFromPrompt(rd *bufio.Reader) string {
	s, _ := rd.ReadString('\n')
	return strings.TrimSuffix(s, "\n")
}
