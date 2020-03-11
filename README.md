# go-telegram

A simple program to send messages to your Telegram account via a bot using CLI.

## Installation

Clone the repository, change API token and chatID that of yours and run:

```sh
cd go-telegram
make build
```

## Usage

go-telegram sends a message to the provided Telegram account ID.

```
Usage: go-telegram [options] <message>

Options:
    --chat-id      Specify the Telegram chat ID to which the message will be sent
    --config       Specify the path to config file
    --help         Displays help message
```

### Examples

```sh
# Sends message to default chat ID
go-telegram "Hello, world!"

# Sends message to a Telegram account with ID 12357851
go-telegram --client-id 12357851 "Hello, world!"

# Sends message to a Telegram account ID 12357851 with the configuration defined at ~/.config/go-telegram  
go-telegram --config ~/.config/go-telegram --client-id 12357851 --send-message "Hello, world!"
```
