package commands

import (
	"io/ioutil"
	"fmt"
	"github.com/zuzuleinen/dave/config"
	"bytes"
)

func Focus() {
	f, err := ioutil.ReadFile("/etc/hosts")

	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	for _, c := range newContent() {
		f = append(f, c)
	}

	ioutil.WriteFile("/etc/hosts", f, 644)
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
