package commands

import (
	"bytes"
	"fmt"
	"github.com/zuzuleinen/dave/config"
	"io/ioutil"
	"strings"
)

const HOSTS_FILE = "/etc/hosts"

func Focus() {
	f, err := ioutil.ReadFile(HOSTS_FILE)

	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	for _, c := range newContent() {
		f = append(f, c)
	}
	ioutil.WriteFile(HOSTS_FILE, f, 644)
}

func ClearFocus() {
	f, err := ioutil.ReadFile(HOSTS_FILE)

	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	final := strings.Replace(string(f), string(newContent()), "", -1)

	ioutil.WriteFile(HOSTS_FILE, []byte(final), 644)
}

func newContent() []byte {
	var b bytes.Buffer

	b.WriteString("\n# Begin Dave focus websites\n")
	for _, website := range config.FavouriteWebSites() {
		b.WriteString("127.0.0.1  " + website + "\n")
	}
	b.WriteString("# End Dave focus websites\n")

	return b.Bytes()
}
