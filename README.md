# Weather CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/agaonker/weather)](https://goreportcard.com/report/github.com/agaonker/weather)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A command-line application to get weather information for any location using the OpenWeatherMap API.

## Repository
```bash
git clone https://github.com/agaonker/weather.git
cd weather
```

## Features

- Get current weather information for any location including:
  - Location name
  - Current temperature
  - Feels like temperature
  - Wind speed
- Supports custom coordinates
- Default location set to San Ramon, CA

## Prerequisites

- Go 1.21 or higher
- OpenWeatherMap API key

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd weather
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file in the project root and add your OpenWeatherMap API key:
```
OPENWEATHER_API_KEY=your_api_key_here
```

## Usage

### Build the application
```bash
go build
```

### Run with default location (San Ramon)
```bash
./weather
```

### Run with custom coordinates
```bash
# Using long flags
./weather --latitude 15.2993 --longitude 74.1240  # For Goa, India

# Using short flags
./weather -a 15.2993 -o 74.1240  # For Goa, India
```

### Get help
```bash
./weather --help
```

## Example Output
```
Weather for Panaji:
-------------------
Current Temperature: 28.5°C
Feels Like: 30.2°C
Wind Speed: 3.2 m/s
-------------------
```

## Getting an API Key

1. Go to [OpenWeatherMap](https://openweathermap.org/)
2. Sign up for a free account
3. Get your API key from your account dashboard
4. Add the API key to your `.env` file

## Project Structure

```
.
├── cmd/
│   └── root.go      # Main command implementation
├── main.go          # Application entry point
├── .env             # Environment variables
├── go.mod           # Go module file
└── README.md        # This file
```

## License

MIT 