package main

import (
	"fmt"
	"os"
	"strconv"

	"flightres/airline"
)

const dataFile = "flights_data.json"

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	dir := airline.Load(dataFile)
	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "schedule":
		if len(args) != 4 {
			fmt.Println("Usage: go run main.go schedule <flightNumber> <origin> <destination> <capacity>")
			return
		}
		capacity, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Capacity must be a number.")
			return
		}
		f := airline.Flight{
			FlightNumber: args[0],
			Origin:       args[1],
			Destination:  args[2],
			Capacity:     capacity,
		}
		if err := dir.ScheduleFlight(f); err != nil {
			fmt.Println("Error:", err)
			return
		}
		dir.Save(dataFile)
		fmt.Println("Scheduled and saved:", f.FlightNumber)

	case "find":
		if len(args) != 1 {
			fmt.Println("Usage: go run main.go find <flightNumber>")
			return
		}
		f, err := dir.FindFlight(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printFlight(*f)

	case "reserve":
		if len(args) != 2 {
			fmt.Println("Usage: go run main.go reserve <flightNumber> <seatCount>")
			return
		}
		count, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Seat count must be a number.")
			return
		}
		if err := dir.ReserveSeat(args[0], count); err != nil {
			fmt.Println("Error:", err)
			return
		}
		dir.Save(dataFile)
		fmt.Println("Reserved and saved:", args[0])

	case "cancel":
		if len(args) != 1 {
			fmt.Println("Usage: go run main.go cancel <flightNumber>")
			return
		}
		if err := dir.CancelFlightRoute(args[0]); err != nil {
			fmt.Println("Error:", err)
			return
		}
		dir.Save(dataFile)
		fmt.Println("Cancelled and saved:", args[0])

	case "list":
		flights := dir.AllFlights()
		if len(flights) == 0 {
			fmt.Println("No flights scheduled.")
			return
		}
		for _, f := range flights {
			printFlight(f)
		}

	default:
		printUsage()
	}
}

func printFlight(f airline.Flight) {
	fmt.Printf("FlightNumber: %s | Origin: %s | Destination: %s | Capacity: %d | Booked: %d\n",
		f.FlightNumber, f.Origin, f.Destination, f.Capacity, f.BookedSeats)
}

func printUsage() {
	fmt.Println("go run main.go schedule <flightNumber> <origin> <destination> <capacity>")
	fmt.Println("go run main.go find     <flightNumber>")
	fmt.Println("go run main.go reserve  <flightNumber> <seatCount>")
	fmt.Println("go run main.go cancel   <flightNumber>")
	fmt.Println("go run main.go list")
	fmt.Println()
	fmt.Println("Flights are saved in", dataFile, "in the current directory.")
}
