package passenger

import (
    "fmt"
    "main/journey/route"
)

type Passenger struct {
    Name string
}

func (p *Passenger) Travel(r *route.Route) {
    fmt.Printf("%s's journey begins.\n", p.Name)
    for i, vehicle := range r.Vehicles {
        fmt.Printf("Step %d: ", i+1)
        vehicle.AcceptPassengers(1)
        vehicle.DropOffPassengers(1)
    }
    fmt.Printf("%s's journey ends.\n", p.Name)
}