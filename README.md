# CLI Personal assistant

## About
Hi! I'm Dave, your personal assistant. I'm not quite ready yet, since my creator has a full-time job and is a lazy bum. But in the future
you will be able to add reminders via CLI which will be sent to your e-mail when the time is due, save credentials and focus at work by blocking websites you frequently visit.

Please note that this is work in progress and my creator is just learning Go. You can always contribute to me by opening as pull request.

## Requirements

Go should be [installed and set up](https://golang.org/doc/install) on your system. Tested with version **go1.7.1**

So far this tool is available only for **Linux** distributions. However besides `focus` and `focus-clear`, the rest of the commands should also work on Windows/Mac OS.

[SQLite](https://sqlite.org/) must be installed on your system. This version was tested with **sqlite3**.

## Installation

`$ go get github.com/zuzuleinen/dave

 $ cd $GOPATH/src/github.com/zuzuleinen/dave/
 
 $ go install
 
 $ dave install
 
 $ dave`


## Usage

```shell
Usage:
    dave <command>

List of commands:
  install:        Install the SQLite database.
  remind:         Create a new reminder.
  reminders:      List all pending reminders.
  credential:     Add a new credential record.
  credentials:    List all credentials.
  focus:          Block websites from config/user.go/FavouriteWebSites
  focus-clear:    Unblock websites from config/user.go/FavouriteWebSites


Options:
  -h --help         Show this screen.
  -v, --version     Show version.
```
Since `dave focus` requires sudo, you might want to alias: ` alias focus='sudo env "PATH=$PATH" dave focus'` and then usage for this command becomes just `focus`
