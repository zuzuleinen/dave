# CLI Personal assistant

## About
Hi! I'm Dave, your personal assistant. I'm not quite ready yet, since my creator has a full-time job and is a lazy bum. But in the future
you will be able to add reminders via CLI which will be sent to your e-mail when the time is due, save credentials and focus at work by blocking websites you frequently visit.

Please note that this is work in progress and my creator is just learning Go. You can always contribute to me by opening as pull request.

## Requirements

Go should be [installed and set up](https://golang.org/doc/install) on your system. Tested with version **go1.7.1**

So far this tool is available only for **Linux** distributions. However besides `focus` and `focus-clear`, the rest of the commands should also work on Windows/Mac OS.

[SQLite](https://sqlite.org/) must be installed on your system. This version was tested with **sqlite3**.

You should have a [mailgun](http://www.mailgun.com/) account. You can set-up a free account there and just use the sandbox credentials. It should be enough for the reminders you add(max 10.000 mails per month).

## Installation

1. Prepare the executable:

```shell
$ go get github.com/zuzuleinen/dave
$ cd $GOPATH/src/github.com/zuzuleinen/dave/
$ go install
$ dave install
$ dave
```
2. Add your e-mail and the list of websites you want to block in *dave/config/user.go*
3. Add mailgun credentials in *dave/config/mailgun.go*


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
