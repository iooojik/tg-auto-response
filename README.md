# Telegram Greeting (auto-hello) Bot

This bot is designed to automatically send a greeting message whenever
a user's message matches predefined messages from a configuration file. Note that this bot is available exclusively for
Telegram Premium users.

## Features

- **Automated Greetings:** Sends a greeting message when a user's message matches the predefined messages in the
  configuration.
- **Configuration-Based:** Easily customizable greetings and trigger messages through a configuration file.
- **Premium Only:** Ensures that the bot interacts only with Telegram Premium users.

## Getting Started

### Prerequisites

- Docker and docker-compose installed
- Telegram Bot API Token (you can get this by creating a bot on [BotFather](https://core.telegram.org/bots#botfather))
- Premium Telegram account for testing and usage

### Configuration

1. Create a `configs/config.yaml` file in the root directory with the structure from `configs/config.yaml`.

2. Replace all values with your desired values.

### Running the Bot

Run the bot:

```bash
docker-compose build && docker-compose up -d
```

## Usage

Once the bot is running, it will automatically monitor messages from users. If a message matches any of the trigger
messages specified in the configuration file, the bot will send the predefined greeting message.

### Example

If your `config.json` contains:

```yaml
bot:
  token: ""
  debug: false
  auto_hello:
  reply: "Привет 👋"
  income_messages:
    - привет
    - hello
    - hi
```

When a user sends a message saying "hello", the bot will respond with "Привет 👋".

---

Enjoy 🎉

