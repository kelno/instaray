# Instaray

[![Latest release](https://img.shields.io/github/v/tag/Madh93/instaray?label=Release)](https://github.com/Madh93/instaray/releases)
[![Go Version](https://img.shields.io/badge/Go-1.23-blue)](https://go.dev/doc/install)
[![Go Reference](https://pkg.go.dev/badge/github.com/Madh93/instaray.svg)](https://pkg.go.dev/github.com/Madh93/instaray)
[![License](https://img.shields.io/badge/License-MIT-brightgreen)](LICENSE)

`Instaray` is a simple [Telegram Bot](https://core.telegram.org/bots) written in [Go](https://go.dev/) that fixes Twitter and Instagram embeds in Telegram using [FxTwitter](https://github.com/FixTweet/FxTwitter) and [InstaFix](https://github.com/Wikidepia/InstaFix).

<p align="center">
  <a href="#features">Features</a> •
  <a href="#requirements">Requirements</a> •
  <a href="#installation">Installation</a> •
  <a href="#Configuration">Configuration</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

## Features

- 🤖 Fixes **twitter.com**, **x.com** and **instagram.com** embeds.
- 🔒 Support **chat ID allowlist**.
- 🐳 **Production-ready Docker image** for easy **deployment**.

## Requirements

- A [Telegram bot token](https://core.telegram.org/bots/features#botfather) (you can get one by talking to [@BotFather](https://t.me/BotFather) on Telegram)

## Installation

### Docker

#### Using `docker run`

Use the `docker run` command to start `Instaray`. Make sure to set the required environment variables:

```sh
docker run --name instaray \
  -e INSTARAY_TELEGRAM_TOKEN=your-telegram-bot-token \
  ghcr.io/madh93/instaray:latest
```

#### Using `docker compose`

Create a `docker-compose.yml` file with the following content:

```yml
services:
  instaray:
    image: ghcr.io/madh93/instaray:latest
    restart: unless-stopped
    # volumes:
    #   - ./custom.config.toml:/var/run/ko/config.default.toml # Optional: specify a custom configuration file instead of the default one
    environment:
      - INSTARAY_TELEGRAM_TOKEN=your-telegram-bot-token
```

Use the `docker compose up` command to start `Instaray`:

```sh
docker compose up
```

### From releases

Download the latest binary from [the releases page](https://github.com/Madh93/instaray/releases):

```sh
curl -L https://github.com/Madh93/instaray/releases/latest/download/toffu_$(uname -s)_$(uname -m).tar.gz | tar -xz -O instaray > /usr/local/bin/instaray
chmod +x /usr/local/bin/instaray
```

### From source

If you have Go installed:

```sh
go install github.com/Madh93/instaray@latest
```

## Configuration

`Instaray` comes with a [default configuration file](config.default.toml) that you can modify to suit your needs.

### Loading a custom configuration file

You can load a different configuration file by using the `-config path/to/config/file` flag when starting the application:

```sh
instaray -config custom.config.tml
```

### Overriding with environment variables

Additionally, you can overridethe configuration values using environment variables that begin with the prefix `INSTARAY_`. This allows you to customize your setup without needing to modify any configuration file:

```sh
INSTARAY_LOGGING_LEVEL=debug INSTARAY_TELEGRAM_ALLOWLIST=chat_id_1,chat_id_2 instaray
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any bug fixes or enhancements.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the [MIT license](LICENSE).
