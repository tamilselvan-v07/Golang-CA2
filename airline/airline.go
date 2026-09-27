package airline

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Flight struct {
	FlightNumber string `json:"flight_number"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Capacity     int    `json:"capacity"`
	BookedSeats  int    `json:"booked_seats"`
}

type Directory struct {
	flights map[string]*Flight
}

func NewDirectory() *Directory {
	return &Directory{flights: make(map[string]*Flight)}
}

func (d *Directory) ScheduleFlight(f Flight) error {
	if f.Origin == f.Destination {
		return errors.New("origin and destination cannot be the same")
	}
	if f.Capacity <= 0 {
		return errors.New("capacity must be greater than zero")
	}
	if _, exists := d.flights[f.FlightNumber]; exists {
		return fmt.Errorf("flight %s already scheduled", f.FlightNumber)
	}
	d.flights[f.FlightNumber] = &f
	return nil
}

func (d *Directory) FindFlight(flightNum string) (*Flight, error) {
	f, ok := d.flights[flightNum]
	if !ok {
		return nil, fmt.Errorf("flight %s not found", flightNum)
	}
	return f, nil
}

func (d *Directory) ReserveSeat(flightNum string, count int) error {
	f, err := d.FindFlight(flightNum)
	if err != nil {
		return err
	}
	if f.BookedSeats+count > f.Capacity {
		return fmt.Errorf("only %d seat(s) left on %s", f.Capacity-f.BookedSeats, flightNum)
	}
	f.BookedSeats += count
	return nil
}

func (d *Directory) CancelFlightRoute(flightNum string) error {
	if _, ok := d.flights[flightNum]; !ok {
		return fmt.Errorf("flight %s not found", flightNum)
	}
	delete(d.flights, flightNum)
	return nil
}

func (d *Directory) AllFlights() []Flight {
	out := make([]Flight, 0, len(d.flights))
	for _, f := range d.flights {
		out = append(out, *f)
	}
	return out
}


func (d *Directory) Save(path string) error {
	b, _ := json.MarshalIndent(d.AllFlights(), "", "  ")
	return os.WriteFile(path, b, 0644)
}

func Load(path string) *Directory {
	d := NewDirectory()
	b, err := os.ReadFile(path)
	if err != nil {
		return d 
	}
	var flights []Flight
	json.Unmarshal(b, &flights)
	for _, f := range flights {
		f := f
		d.flights[f.FlightNumber] = &f
	}
	return d
}