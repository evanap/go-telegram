# go-telegram
A simple program to send messages to your Telegram account via a bot using CLI.


## Installation

Clone the repository, change API token and chatID that of yours and run:
```sh
cd go-telegram
make build
```

## Usage

go-telegram sends a message you supplied using the flag `--send-message`.

```sh
Usage: go-telegram [options] <parameter>

Options:
    --chat-id      Specify the Telegram chat ID to which the message will be sent
    --send-message The message that is going to be sent
    --help         Displays help message 
```