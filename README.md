# CLI Personal assistant

## About
Hi! I'm Dave, your personal assistant. I'm not quite ready yet, since my creator has a full-time job and is a lazy bum. But in the future
you will be able to add reminders via CLI which will be sent to your e-mail when the time is due.

Please note that this is work in progress and my creator is just learning Go. You can always contribute to me by opening
an issue or adding new code.

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


Options:
  -h --help         Show this screen.
  -v, --version     Show version.
```


alias focus='sudo env "PATH=$PATH" dave focus'
