# mudnuts
First Social Sniper Trading(SST) Layer Based on MEV ecosystem

# Mudnuts

## Overview

This project is a social trading bot written in Go, designed to fetch cryptocurrency information and perform socialized sniping trades.

## Features

- Regularly fetches the latest cryptocurrency information.
- Automatically posts cryptocurrency information to a Discord channel.
- Filters and sorts cryptocurrency information.
- Provides on-chain social sniping trade rankings and follow-ups.

## Getting Started

### Prerequisites

- Go 1.15 or higher
- Discord Bot Token

### Installation

1. Clone this repository to your local machine.
2. Compile the project in the /cmd/bot directory using `go build`.
3. Set the environment variables `DISCORD_TOKEN` and `DISCORD_CHANNEL_ID`.

### Usage

Run the compiled executable file, and the bot will start listening and posting information.

## Contributing

Contributions of any kind are welcome. Please open an issue to discuss what you would like to change before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
