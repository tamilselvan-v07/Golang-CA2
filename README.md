# Flight Reservation Directory

A command-line Go project for managing flight routes — schedule flights,
look them up, reserve seats, and cancel routes. Data is saved to a JSON
file on disk, so it survives between runs: close the program, come back
days later, and your flights are still there.

Built as a learning project to practice Go fundamentals: structs, pointer
receivers, maps, error handling, command-line arguments (`os.Args`), and
file persistence with `encoding/json`.

## Domain

**Aviation** — a `Flight` represents one scheduled route between two airports.

## Project Structure

```
Golang-CA1/
├── go.mod
├── main.go              # CLI entry point (parses os.Args)
├── flights_data.json     # created automatically on first save
└── airline/
    └── airline.go        # Flight struct + Directory + JSON persistence
```

## The `Flight` struct

```go
type Flight struct {
    FlightNumber string
    Origin       string
    Destination  string
    Capacity     int
    BookedSeats  int
}
```

## Package API (`airline`)

| Method | Description |
|---|---|
| `ScheduleFlight(f Flight) error` | Adds a new flight. Fails if origin == destination, capacity <= 0, or the flight number is already scheduled. |
| `FindFlight(flightNum string) (*Flight, error)` | Looks up a flight by flight number. Returns an error if not found. |
| `ReserveSeat(flightNum string, count int) error` | Books `count` seats on a flight. Fails if it would exceed capacity. |
| `CancelFlightRoute(flightNum string) error` | Removes a scheduled flight. Fails if the flight number doesn't exist. |
| `AllFlights() []Flight` | Returns every stored flight, for listing. |
| `Save(path string) error` | Writes all flights to a JSON file. |
| `Load(path string) *Directory` | Loads flights from a JSON file (starts empty if the file doesn't exist yet). |

## Getting Started

### Prerequisites

- Go 1.21 or later

### Clone and run

```bash
git clone https://github.com/tamilselvan-v07/Golang-CA1.git
cd Golang-CA1
go run main.go
```

Running with no arguments prints usage help:

```
go run main.go schedule <flightNumber> <origin> <destination> <capacity>
go run main.go find     <flightNumber>
go run main.go reserve  <flightNumber> <seatCount>
go run main.go cancel   <flightNumber>
go run main.go list

Flights are saved in flights_data.json in the current directory.
```

### Example usage

```bash
go run main.go schedule AI202 Chennai Delhi 180
go run main.go schedule 6E345 Mumbai Bangalore 220
go run main.go list
go run main.go reserve AI202 4
go run main.go find AI202
go run main.go cancel 6E345
```

Sample `list` output:

```
FlightNumber: AI202 | Origin: Chennai | Destination: Delhi | Capacity: 180 | Booked: 4
FlightNumber: 6E345 | Origin: Mumbai | Destination: Bangalore | Capacity: 220 | Booked: 0
```

## Data Persistence

- On startup, the program loads `flights_data.json` from the current
  directory if it exists.
- If it doesn't exist yet (first run), it just starts empty — no error.
- After every successful `schedule`, `reserve`, or `cancel`, the entire
  directory is re-saved to that file, so it's never out of sync.
- The file is plain JSON, so it can be opened in a text editor to inspect
  or manually edit.

Example `flights_data.json`:

```json
[
  {
    "flight_number": "AI202",
    "origin": "Chennai",
    "destination": "Delhi",
    "capacity": 180,
    "booked_seats": 4
  }
]
```

## Design Notes

- Flights are stored as `*Flight` (pointers) inside a `map[string]*Flight`,
  so `ReserveSeat` can mutate the stored flight's `BookedSeats` in place
  rather than working on a disconnected copy.
- Errors are returned, not panicked — callers decide how to handle a
  failed reservation or a missing flight.
- Persistence uses only `encoding/json` and `os` from the standard
  library — no external database or driver, so the project runs anywhere
  Go is installed.
- Each CLI invocation runs one command and exits; the JSON file is what
  carries state between separate runs.

## License

MIT
