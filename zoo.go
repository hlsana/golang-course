package main

type Zookeeper struct {
	FirstName  string
	LastName   string
	Occupation string
}

type Animal struct {
	Species  string
	Name     string
	IsInCage bool
}

func main() {

	Lion := Animal{
		Species:  "Lion",
		Name:     "Alex",
		IsInCage: false,
	}

	Zebra := Animal{
		Species:  "Zebra",
		Name:     "Marty",
		IsInCage: false,
	}
	Hippo := Animal{
		Species:  "Hippopotamus",
		Name:     "Gloria",
		IsInCage: false,
	}
	Penguin := Animal{
		Species:  "Penguin",
		Name:     "Skipper",
		IsInCage: false,
	}
	Giraffe := Animal{
		Species:  "Giraffe",
		Name:     "Melman",
		IsInCage: false,
	}

	Zookeeper := Zookeeper{
		FirstName:  "Henry",
		LastName:   "Zoolover",
		Occupation: "Zookeeper",
	}

	if Lion.IsInCage == !true {
		Lion.IsInCage = true
	}
	if Zebra.IsInCage == !true {
		Zebra.IsInCage = true
	}
	if Hippo.IsInCage == !true {
		Hippo.IsInCage = true
	}
	if Penguin.IsInCage == !true {
		Penguin.IsInCage = true
	}
	if Giraffe.IsInCage == !true {
		Giraffe.IsInCage = true
	}

	println(Zookeeper.Occupation, Zookeeper.FirstName, Zookeeper.LastName, "starts catching runaway animals.")
	switch Lion.IsInCage {
	case true:
		println(Lion.Species, Lion.Name, "has been caught.")
	case false:
		println(Lion.Species, Lion.Name, "hasn't been caught.")
	}

	switch Zebra.IsInCage {
	case true:
		println(Zebra.Species, Zebra.Name, "has been caught.")
	case false:
		println(Zebra.Species, Zebra.Name, "hasn't been caught.")
	}

	switch Hippo.IsInCage {
	case true:
		println(Hippo.Species, Hippo.Name, "has been caught.")
	case false:
		println(Hippo.Species, Hippo.Name, "hasn't been caught,")
	}

	switch Penguin.IsInCage {
	case true:
		println(Penguin.Species, Penguin.Name, "has been caught.")
	case false:
		println(Penguin.Species, Penguin.Name, "hasn't been caught.")
	}

	switch Giraffe.IsInCage {
	case true:
		println(Giraffe.Species, Giraffe.Name, "has been caught.")
	case false:
		println(Giraffe.Species, Giraffe.Name, "hasn't been caught.")
	}

}
