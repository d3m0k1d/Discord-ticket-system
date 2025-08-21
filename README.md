# Discord-ticket-system
A universal and configurable ticketing system for Discord with full management via Discord and Telegram.

## Table of contents

1. [Overview](#overview)
2. [Features](#features)
3. [Requirements](#requirements)
4. [Installation](#installation)
5. [License](#license)


## Overview
This project implements a microservices-based ticketing system for Discord.  
All server members can create tickets via the Discord bot, and admins/moderators can manage them both in Discord and Telegram.

## Features(in progress)
- Receive notifications in Telegram for new tickets and updates
- Configurable notification templates, categories, and priorities
- Metrics monitoring (number of tickets, resolution time)
- Full management via Discord and Telegram

## Requirements
  - Docker & Docker Compose
  - Go 1.24+
  - PostgreSQL (v13+)
  - Prometheus (for monitoring)

## Installation

1. Clone the repository
```bash
git clone https://github.com/d3m0k1d/Discord-ticket-system.git
cd Discord-ticket-system
```
2. Build the Docker images
```bash
docker build -t api-gateway -f cmd/api-gateway/Dockerfile .
docker build -t telegram-bot -f cmd/telegram-bot/Dockerfile .
docker build -t discord-bot -f cmd/discord-bot/Dockerfile .
```

3. Start the services with Docker compose
```bash
docker compose up -d
```


## License

This project is licensed under the MIT License. 