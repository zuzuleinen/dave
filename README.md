# Dave - CLI Personal assistant

<p align="center">
<img align="middle" src="dave.jpg" width="200" />
</p>


## About
This is my first project written in Go. Dave is a CLI tool which automates some taks I do at work such as: add a reminder for a certain day and receive and e-mail when time is due, add some credentials from various projects, cancel the noise by blocking certain websites. In the future I might add new commands.


## Requirements

Go should be [installed and set up](https://golang.org/doc/install) on your system. Tested with version **go1.7.1**

So far this tool is available only for **Linux** distributions. However besides `focus` and `focus-clear`, the rest of the commands should also work on Windows/Mac OS.

[SQLite](https://sqlite.org/) must be installed on your system. This version was tested with **sqlite3**.

You should have a [mailgun](http://www.mailgun.com/) account. You can set-up a free account there and just use the sandbox credentials. It should be enough for the reminders you add(max 10.000 mails per month).

## Installation

* Prepare the executable 

```shell
$ go get github.com/zuzuleinen/dave
$ cd $GOPATH/src/github.com/zuzuleinen/dave/
$ go install
$ dave install
$ dave
```
* Add your e-mail and the list of websites you want to block in *dave/config/user.go*
* Add mailgun credentials in *dave/config/mailgun.go*
* Set up the daemon which can be started with `$ dave cli`




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

## Questions or suggestions
If you encounter a problem feel free to [open](https://github.com/zuzuleinen/dave/issues/new) an issue.
