# GoStats

GoStats is a simple system statistics monitoring tool written in Go. It provides real-time information about your system's CPU usage, memory, disk usage, and uptime. The project is modular, allowing you to use it as a command-line tool or extend it for API/server-based monitoring.

## Features

- **CPU Usage**: Displays the current CPU usage percentage.
- **Memory Usage**: Shows used and total memory in GB.
- **Disk Usage**: Shows used and total disk space in GB.
- **Uptime**: Displays how long the system has been running.
- **Modular Design**: Easily extendable for CLI dashboards or API servers.

## Project Structure

```
GoStats/
├── main.go                # Entry point for the application
├── go.mod                 # Go module file
├── collector/
│   └── metrics.go         # System stats collection logic
├── cli/
│   └── dashboard.go       # CLI dashboard (optional/extendable)
├── api/
│   └── server.go          # API server (optional/extendable)
```

## How It Works

- The metrics.go file contains logic to fetch system statistics using the [gopsutil](https://github.com/shirou/gopsutil) library.
- The main.go file demonstrates how to use the collector to print system stats to the console.
- You can extend the project by implementing CLI dashboards or REST APIs using the provided folders.

## Prerequisites

- Go 1.18 or higher (recommended: latest stable)

## Installation

1. **Clone the repository** (if not already):
   ```sh
   git clone https://github.com/AdityaAnandCodes/GoStats.git
   cd GoStats
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

## Usage

### Run as CLI

To display system stats in the terminal:
```sh
# From the project root
# On Windows (PowerShell):
go run main.go
```
You will see output like:
```
System Stats:
CPU Usage: 10.23%
Memory: 8.00 GB / 16.00 GB
Disk: 120.00 GB / 256.00 GB
Uptime: 5h23m12s
```

### Extend as API or Dashboard

- Implement your own CLI dashboard in dashboard.go.
- Build an API server in server.go to serve stats over HTTP.

## Customization

- The metrics.go file can be modified to collect additional system metrics as needed.
- You can add more output formats or integrate with monitoring tools.

## License

This project is licensed under the MIT License.

---
