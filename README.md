# 🌤️ weathergo

A fast, lightweight command-line weather tool built with Go — because checking the weather shouldn't mean leaving your terminal.

---

## 📌 Overview

`weathergo` is a CLI application that fetches real-time weather data for any city in the world and displays it in a clean, readable terminal format. Built as an exploration of Go's core strengths — simplicity, speed, and clean HTTP handling — this project demonstrates idiomatic Go patterns including struct-based JSON parsing, environment-based configuration, and modular code organization.

---

## ⚙️ Tech Stack

| Layer | Technology |
|---|---|
| Language | Go (Golang) |
| Weather Data | OpenWeatherMap API (Current Weather) |
| HTTP Client | `net/http` (Go stdlib) |
| JSON Parsing | `encoding/json` (Go stdlib) |
| Env Management | `godotenv` |
| Version Control | Git + GitHub |

---

## 🗂️ Project Structure

```
weathergo/
├── main.go        # Entry point — handles CLI args and orchestration
├── weather.go     # API call, JSON struct definitions, data fetching
├── display.go     # Terminal output formatting and emoji mapping
├── .env           # Local API key storage (git-ignored)
├── .gitignore     # Excludes .env and build artifacts
└── go.mod         # Go module definition
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed
- A free API key from [OpenWeatherMap](https://openweathermap.org/api)

### Installation

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/weathergo.git

# Navigate into the project
cd weathergo

# Install dependencies
go mod tidy
```

### Configuration

Create a `.env` file in the project root:

```
WEATHER_API_KEY=your_api_key_here
```

> ⚠️ Never commit your `.env` file. It is already included in `.gitignore`.

### Run

```bash
go run . <city>
```

**Example:**
```bash
go run . Bengaluru
```

### Build a Binary

```bash
go build -o weathergo .
./weathergo Bengaluru
```

---

## 📸 Sample Output

```
─────────────────────────────────
  📍 Bengaluru, IN
─────────────────────────────────
  ☁️  Scattered Clouds

  🌡️  Temperature  : 27.3°C
  🤔  Feels Like   : 29.1°C
  💧  Humidity     : 68%
  💨  Wind Speed   : 12.4 km/h
─────────────────────────────────
```

---

## 🧠 Key Concepts Demonstrated

- **Idiomatic Go project structure** — separation of concerns across multiple `.go` files in a single package
- **Struct-based JSON unmarshalling** — mapping nested API responses to typed Go structs
- **CLI argument handling** — using `os.Args` for a zero-dependency command-line interface
- **Environment variable management** — secure API key handling via `.env` and `godotenv`
- **Error handling** — Go-style explicit error propagation with descriptive messages
- **HTTP client usage** — making and handling REST API calls using Go's stdlib

---

## 🔮 Potential Enhancements

- [ ] Support for multi-day forecast
- [ ] Add `--unit` flag to toggle Celsius / Fahrenheit
- [ ] Coloured terminal output using `fatih/color`
- [ ] Support for multiple cities in a single command
- [ ] Cache last result locally to support offline mode

---

## 👨‍💻 Author

**Sarthak Sengupta**  
Software Developer | Python · Go · AWS · Automation  
[LinkedIn](https://linkedin.com/in/sarthak-sengupta) · [GitHub](https://github.com/YOUR_USERNAME)

---

## 📄 License

MIT License — feel free to fork, modify, and build on this.
