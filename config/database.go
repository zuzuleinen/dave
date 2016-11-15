package config

import (
	"github.com/mitchellh/go-homedir"
	"os"
)

const DB_FILE = "davedb.db"

func TimeFormat() string {
	return "Monday, 2 Jan 2006 at 15:04"
}

func DbPath() string {
	homeDir, _ := homedir.Dir()
	return homeDir + string(os.PathSeparator) + DB_FILE
}
