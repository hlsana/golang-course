package route

import (
    "fmt"
    "main/journey/transport"
)

type Route struct {
    Vehicles []transport.PublicTransport
}

func (r *Route) AddVehicleToRoute(vehicle transport.PublicTransport) {
    r.Vehicles = append(r.Vehicles, vehicle)
    fmt.Println("Added 1 vehicle to route.")
}

func (r *Route) ShowListOfVehiclesOnRoute() {
    fmt.Println("Vehicles on route:")
    for i, vehicle := range r.Vehicles {
        fmt.Printf("%d. %T\n", i+1, vehicle)
    }
}