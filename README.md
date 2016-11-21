# Dave - CLI Personal assistant

<p align="center">
<img align="middle" src="dave.jpg" width="200" />
</p>


## About
This is my first project written in Go. Dave is a CLI tool which automates some taks I do at work such as: add a reminder for a certain day and receive and email me when time is due, add credentials from various projects, cancel the noise by blocking certain websites. In the future I might add new commands.


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
    dave <command> [argument]
    dave  check    'a link to an article you want to read later'

List of commands:
  install:        Install the SQLite database.
  remind:         Create a new reminder.
  reminders:      List all pending reminders.
  credential:     Add a new credential record.
  credentials:    List all credentials.
  focus:          Block websites from config/user.go/FavouriteWebSites
  focus-clear:    Unblock websites from config/user.go/FavouriteWebSites
  check:          Send an e-mail with subject "check this" and body the string argument



Options:
  -h --help         Show this screen.
  -v, --version     Show version.
```

Since `dave focus` and `dave focus-clear` requires sudo, you should add an alias. Add these 2 lines in you .bashrc file:
` alias focus='sudo env "PATH=$PATH" dave focus'`<br>
` alias focus-clear='sudo env "PATH=$PATH" dave focus-clear'`</br>
Now, you can use `focus` and `focus-clear` commands.

## Setting up the daemon
To be able to receive your reminders, you need to set up the daemon which basically just loops over you reminders and send them to your e-mail when it's time.
Since I'm using Ubuntu, I use [upstart](http://upstart.ubuntu.com/). But you can also use [Supervisor](http://supervisord.org/) or [daemonize](http://software.clapper.org/daemonize/) to set up the daemon.
Here are the steps for setting up your daemon with upstart:

Create a configuration file for your daemon:
`sudo vim /etc/init/dave.conf`

Add the following contents and make sure to replace andrei with your user and the correct path to dave executable:
```
description "Run the dave daemon"

setuid andrei
start on runlevel [2345]

exec /home/andrei/Projects/bin/dave cli
```
Check that your syntax is OK:
```
sudo init-checkconf /etc/init/dave.conf
```
Start your service
````
sudo service dave start
```

## Questions or suggestions
If you encounter a problem feel free to [open](https://github.com/zuzuleinen/dave/issues/new) an issue.
