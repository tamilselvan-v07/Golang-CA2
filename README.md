# Flight Reservation Directory

A Go project with an interactive command-line interface for managing flight
routes — scheduling flights, looking them up, reserving seats, and
cancelling routes. Data is persisted to a JSON file on disk, so it survives
between runs (close the program today, reopen it in two days, and your
flights are still there).

Built as a learning project to practice Go fundamentals: structs, pointer
receivers, maps, idiomatic error handling, CLI input with `bufio`, and
file I/O with `encoding/json`.

## Domain

**Aviation** — a `Flight` represents one scheduled route between two airports.

## Project Structure

```
Golang-CA1/
├── go.mod
├── main.go              # interactive CLI menu
├── flights_data.json     # created automatically on first save
└── airline/
    └── airline.go        # Flight struct + Directory + file persistence
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

## API

All operations live on a `*Directory`, created with `airline.NewDirectory()`.

| Method | Description |
|---|---|
| `ScheduleFlight(f Flight) error` | Adds a new flight. Fails if origin == destination, capacity <= 0, flight number is empty, or the flight number is already scheduled. |
| `FindFlight(flightNum string) (*Flight, error)` | Looks up a flight by flight number. Returns an error if not found. |
| `ReserveSeat(flightNum string, count int) error` | Books `count` seats on a flight. Fails if the count is non-positive or would exceed capacity. |
| `CancelFlightRoute(flightNum string) error` | Removes a scheduled flight. Fails if the flight number doesn't exist. |

## Getting Started

### Prerequisites

- Go 1.21 or later

### Run it

```bash
git clone https://github.com/tamilselvan-v07/Golang-CA1.git
cd Golang-CA1
go run main.go
```

You'll see a menu:

```
===== Flight Reservation Directory =====
1. Schedule a new flight
2. Find a flight
3. Reserve seat(s)
4. Cancel a flight route
5. List all flights
6. Exit
Choose an option:
```

Every time you schedule, reserve, or cancel, the program immediately writes
the current state to `flights_data.json` in the project folder. Exit with
option 6 (or Ctrl+C — the file is already up to date from the last action),
and next time you run `go run main.go` your flights will still be loaded.

### Data persistence

- On startup, the program loads `flights_data.json` if it exists.
- If it doesn't exist yet (first run), it just starts empty — no error.
- After every successful schedule / reserve / cancel, the whole directory
  is re-saved to that file, so it never falls out of sync.
- The file is plain JSON, so you can open it in a text editor to inspect
  or manually tweak it if needed.

## Design Notes

- Flights are stored as `*Flight` (pointers) inside a `map[string]*Flight`,
  so that `ReserveSeat` can mutate the stored flight's `BookedSeats` in
  place rather than working on a disconnected copy.
- Errors are returned, not panicked — callers decide how to handle a
  failed reservation or a missing flight.
- Persistence uses `encoding/json` + `os.WriteFile` / `os.ReadFile` — no
  external database or driver needed, which keeps the project dependency-free
  and easy to run anywhere Go is installed.

## License

MIT
