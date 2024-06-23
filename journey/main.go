package main

import (
	"fmt"
    "main/journey/route"
    "main/journey/transport"
	"main/journey/passenger"
)

func main() {
	fmt.Println("Who is travelling?")
	traveler := &passenger.Passenger{}

	fmt.Scan(&traveler.Name)

    bus := &transport.Bus{}
    train := &transport.Train{}
    airplane := &transport.Airplane{}

    journeyRoute := &route.Route{}
    journeyRoute.AddVehicleToRoute(bus)
    journeyRoute.AddVehicleToRoute(train)
    journeyRoute.AddVehicleToRoute(airplane)

    journeyRoute.ShowListOfVehiclesOnRoute()

    traveler.Travel(journeyRoute)
}
