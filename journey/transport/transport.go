package transport

import "fmt"

type PublicTransport interface {
    AcceptPassengers(count int)
    DropOffPassengers(count int)
}

type Airplane struct {
    Passengers int
}

func (a *Airplane) AcceptPassengers(count int) {
    a.Passengers += count
    fmt.Printf("Airplane accepted %d passengers, now has %d passengers.\n", count, a.Passengers)
}

func (a *Airplane) DropOffPassengers(count int) {
    if count > a.Passengers {
        count = a.Passengers
    }
    a.Passengers -= count
    fmt.Printf("Airplane dropped off %d passengers, now has %d passengers.\n", count, a.Passengers)
}

type Bus struct {
    Passengers int
}

func (b *Bus) AcceptPassengers(count int) {
    b.Passengers += count
    fmt.Printf("Bus accepted %d passengers, now has %d passengers.\n", count, b.Passengers)
}

func (b *Bus) DropOffPassengers(count int) {
    if count > b.Passengers {
        count = b.Passengers
    }
    b.Passengers -= count
    fmt.Printf("Bus dropped off %d passengers, now has %d passengers.\n", count, b.Passengers)
}
type Train struct {
    Passengers int
}

func (t *Train) AcceptPassengers(count int) {
    t.Passengers += count
    fmt.Printf("Train accepted %d passengers, now has %d passengers.\n", count, t.Passengers)
}

func (t *Train) DropOffPassengers(count int) {
    if count > t.Passengers {
        count = t.Passengers
    }
    t.Passengers -= count
    fmt.Printf("Train dropped off %d passengers, now has %d passengers.\n", count, t.Passengers)
}