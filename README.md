# 🚀 Full Cycle Challenge: Temperature by Zip Code (CEP)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Status](https://img.shields.io/badge/Status-Completed-success?style=flat)]()

This project is a microservice implemented in Go that receives a Brazilian ZIP code (CEP), identifies the corresponding city via **ViaCEP**, and returns the current temperature in **Celsius**, **Fahrenheit**, and **Kelvin** using the **WeatherAPI**.

## 🧠 Architecture & Logic

The system is designed following clean principles to ensure clear separation of concerns:

1.  **Validation**: Strict 8-digit ZIP code validation before any external call.
2.  **Location Fetching**: Integration with **ViaCEP** to resolve the city name.
3.  **Weather Fetching**: Integration with **WeatherAPI** using the city name to get the current temperature in Celsius.
4.  **Conversion Engine**: 
    -   **Fahrenheit**: `C * 1.8 + 32`
    -   **Kelvin**: `C + 273.15` (following standard scientific conversion).

---

## 🚀 Cloud Run URL
A aplicação está implantada e acessível em:
**URL Base:** `https://go-temperature-by-cep-challenge-790564571764.us-central1.run.app/{cep}`

**Exemplo de teste (CEP de São Paulo):**
[https://go-temperature-by-cep-challenge-790564571764.us-central1.run.app/01153000](https://go-temperature-by-cep-challenge-790564571764.us-central1.run.app/01153000)

---

## 📁 Project Structure

Following the workspace's established pattern:

```text
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── entity/          # Business logic (Weather struct and conversions)
│   ├── usecase/         # Application orchestration (GetWeather logic)
│   └── infra/           
│       ├── web/         # HTTP Handler and Router
│       ├── viacep/      # ViaCEP API Client implementation
│       └── weatherapi/  # WeatherAPI Client implementation
├── Dockerfile           # Multi-stage build for efficient deployment
└── README.md            # Documentation
```

---

## ⚙️ Configuration

The application requires a **WeatherAPI** key. You can get one for free at [weatherapi.com](https://www.weatherapi.com/).

Manage your settings via environment variables (or a `.env` file for local development):

| Variable | Description |
|----------|-------------|
| `WEATHER_API_KEY` | API Key for WeatherAPI |
| `PORT` | Port for the web server (default: 8080) |

---

## 🚀 How to Run

### 1. Run with Docker
```bash
docker build -t go-weather .
docker run -p 8080:8080 -e WEATHER_API_KEY=your_key_here go-weather
```

### 2. Run Locally
```bash
export WEATHER_API_KEY=your_key_here
go run cmd/server/main.go
```

### 3. Running Tests
```bash
go test ./...
```

## 🛡️ Rate Limiting
To ensure service availability and stay within the Free Tier limits, the API implements a rate limit:
- **Limit**: 10 requests per minute per IP.
- **Exceeded Limit**: Returns HTTP `429 Too Many Requests`.

---

## 📊 API Usage

### Success Scenario
**Request:**
`GET /01153000`

**Response (200 OK):**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Invalid Zip Code
**Request:**
`GET /123`

**Response (422 Unprocessable Entity):**
`invalid zipcode`

### Zip Code Not Found
**Request:**
`GET /99999999`

**Response (404 Not Found):**
`can not find zipcode`

---

## 🛠️ Technologies
- **Go** (Golang)
- **Chi Router**
- **Docker**
- **Google Cloud Run**
- **ViaCEP API**
- **WeatherAPI**
